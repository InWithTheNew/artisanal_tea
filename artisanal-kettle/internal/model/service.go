package model

// ServiceModel defines the interface for a service, providing methods to access its properties.
type ServiceModel interface {
	GetName() string          // Returns the name of the service
	GetServer() string        // Returns the server address of the service
	GetKubernetesCheck() bool // Returns true if the service is on Kubernetes
}

// Service is a concrete implementation of ServiceModel, representing a deployable service.
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

// Compile-time check to ensure Service implements ServiceModel.
var _ ServiceModel = (*Service)(nil)
