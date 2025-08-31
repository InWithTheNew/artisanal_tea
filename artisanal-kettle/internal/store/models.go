package store

import (
	"artisanal-kettle/internal/model"
	"context"

	"github.com/redis/go-redis/v9"
)

// ServiceStore defines the interface for saving and retrieving services from a store.
type ServiceStore interface {
	SaveService(ctx context.Context, s model.ServiceModel) error               // Save a service to the store
	GetService(ctx context.Context, name string, out model.ServiceModel) error // Retrieve a service by name
	GetAllServices(ctx context.Context) ([]model.ServiceModel, error)          // Retrieve all services
}

// RedisServiceStore implements ServiceStore using Redis as the backend.
type RedisServiceStore struct {
	client *redis.Client // Redis client instance
}
