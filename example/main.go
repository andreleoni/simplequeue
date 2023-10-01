package main

import (
	"fmt"
	"simplequeue"
	"simplequeue/queue/redis"
	"time"
)

func main() {
	workerA := WorkerA{}
	workerA.SetOptions(simplequeue.NewWorkerOptions(0, "default", "WorkerA"))
	simplequeue.Register(workerA.StructName(), &workerA)

	workerB := WorkerB{}
	workerB.SetOptions(simplequeue.NewWorkerOptions(0, "low", "WorkerB"))
	simplequeue.Register(workerB.StructName(), &workerB)

	go func() {
		for {
			// Generate messages to test consumer
			for i := 0; i < 1; i++ {
				for _, worker := range simplequeue.Registers {
					go worker.PerformAsync(fmt.Sprint("my message ", time.Now()))
				}
			}

			time.Sleep(3 * time.Second)
		}
	}()

	consumer := simplequeue.NewConsumer(
		redis.NewRedisQueue(),
		simplequeue.ConsumerOpts{QueuesName: []string{"default", "low", "high"}},
	)

	err := consumer.Consume()

	fmt.Println("error got on server start", err)
}

type WorkerA struct {
	simplequeue.WorkerBase
}

func (w WorkerA) Perform(data string) error {
	fmt.Println("processing worker A", data)

	return nil
}

type WorkerB struct {
	simplequeue.WorkerBase
}

func (w WorkerB) Perform(data string) error {
	fmt.Println("processing worker B", data)

	return nil
}
