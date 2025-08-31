package store

import (
	"artisanal-kettle/internal/model"
	"context"
)

// ListServices retrieves all services from the Redis store and returns them as a slice of ServiceModel.
func ListServices() ([]model.ServiceModel, error) {
	var services []model.ServiceModel
	redis := GetStoreClient()                                 // Get the global Redis client
	rss := NewServiceStore(redis)                             // Create a new RedisServiceStore
	services, err := rss.GetAllServices(context.Background()) // Fetch all services
	if err != nil {
		return nil, err // Return error if retrieval fails
	}
	return services, nil // Return the list of services
}
