package store

import (
	"context"
	"encoding/json"
	"fmt"

	"artisanal-kettle/internal/model"

	"github.com/redis/go-redis/v9"
)

// NewServiceStore creates a new RedisServiceStore with the given Redis client.
func NewServiceStore(client *redis.Client) *RedisServiceStore {
	return &RedisServiceStore{client: client}
}

var redisClient *redis.Client // Global Redis client

// InitRedis initializes the global Redis client with the given connection parameters.
func InitRedis(addr, password string, db int) {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})
}

// GetStoreClient returns the global Redis client instance.
func GetStoreClient() *redis.Client {
	return redisClient
}

// SaveService marshals a ServiceModel and saves it to Redis under its name as the key.
func (r *RedisServiceStore) SaveService(ctx context.Context, s model.ServiceModel) error {
	data, err := json.Marshal(s)
	if err != nil {
		return err
	}
	return r.client.Set(ctx, s.GetName(), data, 0).Err()
}

// Deletes service from backend store.
func (r *RedisServiceStore) DeleteService(ctx context.Context, s model.ServiceModel) error {
	return r.client.Del(ctx, s.GetName()).Err()
}

// GetService retrieves a service by name and unmarshals it into the provided output struct.
// Note: The caller must pass a pointer to the correct concrete type for unmarshalling
func (r *RedisServiceStore) GetService(ctx context.Context, name string, out model.ServiceModel) error {
	val, err := r.client.Get(ctx, name).Result()
	if err != nil {
		return err
	}
	return json.Unmarshal([]byte(val), out)
}

// GetAllServices retrieves all services from Redis and returns them as a slice of ServiceModel.
func (r *RedisServiceStore) GetAllServices(ctx context.Context) ([]model.ServiceModel, error) {
	var services []model.ServiceModel

	// Use KEYS or SCAN to get all keys (for demo; SCAN is better for production)
	keys, err := r.client.Keys(ctx, "*").Result()
	if err != nil {
		return nil, err
	}

	for _, key := range keys {
		val, err := r.client.Get(ctx, key).Result()
		if err != nil {
			continue // skip missing/bad keys
		}
		// Unmarshal into concrete type *model.Service
		svc := &model.Service{}
		err = json.Unmarshal([]byte(val), svc)
		if err != nil {
			continue // skip bad data
		}
		services = append(services, svc)
	}
	fmt.Print(services) // Debug print of all loaded services
	return services, nil
}
