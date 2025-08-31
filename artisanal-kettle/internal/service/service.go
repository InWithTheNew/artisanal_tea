package service

import (
	"artisanal-kettle/internal/command"
	"artisanal-kettle/internal/model"
	"artisanal-kettle/internal/store"
	"context"
	"log"
)

// Service is a concrete implementation of model.ServiceModel, representing a deployable service.
type Service struct {
	Name         string `json:"name"`         // The name of the service
	Server       string `json:"server"`       // The server address for the service
	IsKubernetes bool   `json:"isKubernetes"` // Whether the service is deployed on Kubernetes
}

// GetName returns the service's name.
func (s *Service) GetName() string { return s.Name }

// GetServer returns the service's server address.
func (s *Service) GetServer() string { return s.Server }

// GetKubernetesCheck returns true if the service is on Kubernetes.
func (s *Service) GetKubernetesCheck() bool { return s.IsKubernetes }

// Compile-time check to ensure Service implements model.ServiceModel.
var _ model.ServiceModel = (*Service)(nil)

// New creates a new empty Service instance.
func New() *Service {
	return &Service{}
}

// SubmitNewServiceConfig saves the current service configuration to the Redis store.
func (s *Service) SubmitNewServiceConfig() error {
	ctx := context.Background()
	state := store.GetStoreClient()
	ctn := store.NewServiceStore(state)

	err := ctn.SaveService(ctx, s)

	if err != nil {
		return err
	}

	log.Printf("new service submitted: %s", s.Name)

	return nil
}

func (s *Service) DeleteServiceConfig() error {
	ctx := context.Background()
	state := store.GetStoreClient()
	ctn := store.NewServiceStore(state)

	err := ctn.DeleteService(ctx, s)
	if err != nil {
		return err
	}

	log.Printf("service deleted: %s", s.Name)

	return nil
}

// SubmitCommand retrieves a service by name and sends a command to it using the action package.
func SubmitCommand(svcName, cmd string) (string, error) {

	svc, err := GetService(svcName)
	if err != nil {
		return "", err
	}

	resp, err := command.Send(svc, cmd)

	if err != nil {
		return "", err
	}

	log.Printf("command send to service %s: %s", svcName, cmd)

	return resp, nil
}

// GetService retrieves a service from the store by its name.
func GetService(svcName string) (*Service, error) {

	ctx := context.Background()
	state := store.GetStoreClient()
	r := store.NewServiceStore(state)

	svc := &Service{}
	err := r.GetService(ctx, svcName, svc)
	if err != nil {
		return nil, err
	}
	return svc, nil
}
