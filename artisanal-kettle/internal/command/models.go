package command

// StoreService is a legacy struct for services with nested environment info.
type StoreService struct {
	ServiceName string `json:"serviceName"` // Name of the service
	Environment struct {
		Name         string `json:"name"`         // Environment name
		Server       string `json:"server"`       // Server address
		IsKubernetes bool   `json:"isKubernetes"` // Is this a Kubernetes service?
	} `json:"environment"`
}

// Service is a flat struct for a service definition.
type Service struct {
	Name         string `json:"name"`         // Name of the service
	Server       string `json:"server"`       // Server address
	IsKubernetes bool   `json:"isKubernetes"` // Is this a Kubernetes service?
}

// ServiceModel defines the interface for a service, used for polymorphism.
type ServiceModel interface {
	GetKubernetesCheck() bool // Returns true if the service is on Kubernetes
	GetName() string          // Returns the service's name
	GetServer() string        // Returns the server address
}

// GetName returns the service's name.
func (s *Service) GetName() string { return s.Name }

// GetServer returns the service's server address.
func (s *Service) GetServer() string { return s.Server }

// GetKubernetesCheck returns true if the service is on Kubernetes.
func (s *Service) GetKubernetesCheck() bool { return s.IsKubernetes }
