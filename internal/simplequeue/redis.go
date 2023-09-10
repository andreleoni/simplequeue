package simplequeue

import (
	ctx "context"

	"github.com/redis/go-redis/v9"
)

var RedisAdapter = RedisRepository{
	redisClient: RedisClient(),
}

type RedisRepository struct {
	redisClient *redis.Client
}

func RedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func (r RedisRepository) Push(queueName, data string) error {
	return r.redisClient.LPush(ctx.Background(), queueName, data).Err()
}

func (r RedisRepository) Pop(queueName string, count int) ([]string, error) {
	return r.redisClient.RPopCount(ctx.Background(), queueName, count).Result()
}
