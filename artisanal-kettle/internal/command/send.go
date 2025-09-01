package command

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/ssh"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/remotecommand"
)

// Send dispatches a command to a service, using either SSH or Kubernetes exec depending on the service type.
func Send(svc ServiceModel, cmd, submittingUser string) (string, error) {

	var response string
	var err error

	cmd = sanitiseCommand(cmd)
	log.Printf("%s sent command to %s: %s. Host: %s", submittingUser, svc.GetName(), cmd, svc.GetServer())

	switch svc.GetKubernetesCheck() {
	case true:
		// If the service is on Kubernetes, exec the command in a pod
		response, err = SendArtisanCommandToKubernetes(svc.GetServer(), cmd)
	case false:
		// Otherwise, use SSH to run the command on the server
		response, err = SendArtisanCommandToServer(svc.GetServer(), cmd)
	}

	return response, err
}

// SendArtisanCommandToKubernetes finds a running pod in any namespace and execs the command, returning the output.
func SendArtisanCommandToKubernetes(host string, cmd string) (string, error) {
	// Use in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		return "", fmt.Errorf("failed to get in-cluster config: %w", err)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return "", fmt.Errorf("failed to create k8s client: %w", err)
	}

	// List all namespaces
	nsList, err := clientset.CoreV1().Namespaces().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return "", fmt.Errorf("failed to list namespaces: %w", err)
	}

	for _, ns := range nsList.Items {
		// List all pods in the namespace
		pods, err := clientset.CoreV1().Pods(ns.Name).List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			continue
		}
		for _, pod := range pods.Items {
			if pod.Status.Phase == corev1.PodRunning {
				// Exec into the first running pod found
				return execInPod(clientset, config, ns.Name, pod.Name, cmd)
			}
		}
	}
	return "", fmt.Errorf("no running pods found in any namespace")
}

// execInPod runs a command in the specified pod and returns the output.
func execInPod(clientset *kubernetes.Clientset, config *rest.Config, namespace, podName, command string) (string, error) {
	req := clientset.CoreV1().RESTClient().Post().
		Resource("pods").
		Name(podName).
		Namespace(namespace).
		SubResource("exec").
		Param("container", "").
		Param("command", "/bin/sh").
		Param("command", "-c").
		Param("command", command).
		Param("stdin", "false").
		Param("stdout", "true").
		Param("stderr", "true").
		Param("tty", "false")

	exec, err := remotecommand.NewSPDYExecutor(config, "POST", req.URL())
	if err != nil {
		return "", err
	}
	var stdout, stderr bytes.Buffer
	err = exec.StreamWithContext(context.TODO(), remotecommand.StreamOptions{
		Stdout: &stdout,
		Stderr: &stderr,
	})
	if err != nil {
		return "", fmt.Errorf("exec error: %v, stderr: %s", err, stderr.String())
	}
	return stdout.String(), nil
}

// SendArtisanCommandToServer runs the command on a remote server using SSH and returns the output.
func SendArtisanCommandToServer(host string, cmd string) (string, error) {
	user := os.Getenv("SSH_USER")
	key := os.Getenv("SSH_KEY")

	key = strings.ReplaceAll(key, `\n`, "\n")

	signer, err := ssh.ParsePrivateKey([]byte(key))
	if err != nil {
		return "", err
	}

	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.PublicKeys(signer),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", host+":22", config)
	if err != nil {
		return "", err
	}
	defer client.Close()
	session, err := client.NewSession()
	if err != nil {
		return "", err
	}
	defer session.Close()

	output, err := session.CombinedOutput("php artisan " + cmd)
	return string(output), err
}

// Try to format the cmd input. Prevent cmd injection, add or dedup 'php artisan'.
func sanitiseCommand(cmd string) string {

	strings.ToLower(cmd)
	// Some attempt at preventing cmd injection.
	cleansedCommand := strings.Split(cmd, ";")[0]
	cleansedCommand = strings.Split(cleansedCommand, "&")[0]
	cleansedCommand = strings.Split(cleansedCommand, "|")[0]

	if strings.Contains(cleansedCommand, "php artisan") {
		cleansedCommand = strings.Replace(cleansedCommand, "php artisan ", "", -1)
	}

	return cleansedCommand
}
