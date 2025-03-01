package redis_service

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

var client *redis.Client
var ctx context.Context

func Start() {

	// Create a new context
	ctx = context.Background()

	// Connect to Redis
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // Address of the Redis server
		Password: "",               // No password set
		DB:       0,                // Use default DB
	})

	// Ping the Redis server to check if the connection is successful
	_, err := client.Ping(ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}

	// Clean up and close the Redis client
	//defer client.Close()
}

func Set(key string, value interface{}, expiration time.Duration) error {
	// Set a key-value pair in Redis
	err := client.Set(ctx, key, value, expiration).Err()
	return err
	//if err != nil {
	//	log.Fatalf("Could not set key: %v", err)
	//}
}

func Get(key string) (interface{}, error) {

	val, err := client.Get(ctx, key).Result()
	if err != nil {
		//log.Printf("Could not get key: %v", err)
		return nil, err
	}

	return val, nil
}
