package store

import (
	"artisanal-kettle/internal/model"
	"context"
)

// ServiceModel interface is duplicated here for legacy reasons; prefer using model.ServiceModel.
type ServiceModel interface {
	GetName() string          // Returns the name of the service
	GetServer() string        // Returns the server address of the service
	GetKubernetesCheck() bool // Returns true if the service is on Kubernetes
}

// ListServices retrieves all services from the Redis store and returns them as a slice of ServiceModel.
func ListServices() ([]model.ServiceModel, error) {
	var services []model.ServiceModel
	redis := GetRedisClient()                                 // Get the global Redis client
	rss := NewServiceStore(redis)                             // Create a new RedisServiceStore
	services, err := rss.GetAllServices(context.Background()) // Fetch all services
	if err != nil {
		return nil, err // Return error if retrieval fails
	}
	return services, nil // Return the list of services
}
