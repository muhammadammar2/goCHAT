package redisclient

import (
	"context"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

// Create a new Redis client
func NewRedisClient() *redis.Client {
    client := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379", // Redis server address
        Password: "",               // No password
        DB:       0,                // Use default DB
    })

    // Test the connection
    _, err := client.Ping(ctx).Result()
    if err != nil {
        log.Fatalf("Could not connect to Redis: %v", err)
    } else {
        log.Println("Successfully connected to Redis")
    }

    return client
}

// Function to blacklist a token
func BlacklistToken(client *redis.Client, token string, expiration time.Duration) error {
    err := client.Set(ctx, token, "blacklisted", expiration).Err()
    if err != nil {
        log.Printf("Error blacklisting token: %v", err)
        return err
    }
    log.Printf("Token blacklisted successfully: %s", token)
    return nil
}

// Function to check if token is blacklisted
func IsTokenBlacklisted(client *redis.Client, token string) (bool, error) {
    result, err := client.Get(ctx, token).Result()
    if err == redis.Nil {
        log.Printf("Token not found in blacklist: %s", token)
        return false, nil // Token is not blacklisted
    } else if err != nil {
        log.Printf("Error checking token blacklist: %v", err)
        return false, err // Redis error
    }
    isBlacklisted := result == "blacklisted"
    log.Printf("Token blacklist check: %s, Blacklisted: %v", token, isBlacklisted)
    return isBlacklisted, nil
}