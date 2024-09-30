package redis

import (
	"context"
	"github.com.ivanrafli14/ecommerce-golang/pkg/config"
	"github.com.ivanrafli14/ecommerce-golang/pkg/response"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

type Interface interface {
	SetData(key string, value bool, ttl time.Duration) error
	GetData(key string) (any, error)
}
type redisClient struct {
	Client *redis.Client
}

func NewRedis(cfg config.RedisConfig) Interface {
	username := cfg.Username
	password := cfg.Password

	redisNewClient := redis.NewClient(&redis.Options{
		Addr:     username,
		Password: password,
	})

	err := redisNewClient.Ping(context.Background()).Err()
	if err != nil {

		log.Fatal(err)
	}

	return &redisClient{
		Client: redisNewClient,
	}
}

func (r *redisClient) SetData(key string, value bool, ttl time.Duration) error {
	ctx := context.Background()
	err := r.Client.Set(ctx, key, value, ttl).Err()

	if err != nil {
		log.Println("Error in redis: ", err)
		return err
	}
	return nil
}

func (r *redisClient) GetData(key string) (any, error) {
	val, err := r.Client.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return nil, response.ErrJWTExpired
	} else if err != nil {
		log.Println(err)
		return nil, err
	}
	return val, nil
}
