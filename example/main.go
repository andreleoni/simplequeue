package main

import (
	"fmt"
	"simplequeue"
	"simplequeue/example/internal/workers"
	_ "simplequeue/example/internal/workers"
	"simplequeue/queue/redis"
	"time"
)

func main() {
	generateSeedToQueues()

	consumer := simplequeue.NewConsumer(
		redis.NewRedisQueue(),
		simplequeue.ConsumerOpts{QueuesName: []string{"default", "low", "high"}},
	)

	err := consumer.Consume()

	fmt.Println("error got on server start", err)
}

func generateSeedToQueues() {
	go func() {
		for {
			// Generate messages to test consumer
			for i := 0; i < 1; i++ {
				go workers.NewWorkerA().PerformAsync(fmt.Sprint("my message A ", time.Now()))
				go workers.NewWorkerB().PerformAsync(fmt.Sprint("my message B ", time.Now()))
			}

			time.Sleep(3 * time.Second)
		}
	}()
}
