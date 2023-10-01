package redis

import (
	ctx "context"

	"github.com/redis/go-redis/v9"
)

type RedisQueue struct {
	client *redis.Client
}

func NewRedisQueue() RedisQueue {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	redisQueue := RedisQueue{client: redisClient}

	return redisQueue
}

func (r RedisQueue) Push(queueName, data string) error {
	return r.client.LPush(ctx.Background(), queueName, data).Err()
}

func (r RedisQueue) Pop(queueName string, count int) ([]string, error) {
	return r.client.RPopCount(ctx.Background(), queueName, count).Result()
}
