package simplequeue

import (
	"encoding/json"
	"log"
	"simplequeue/queue"
	"simplequeue/queue/redis"
)

type Enqueuer struct {
	Queue queue.Queue
}

func NewEnqueuer() Enqueuer {
	return Enqueuer{
		Queue: redis.NewRedisQueue(),
	}
}

func (e Enqueuer) Enqueue(queue, workerName, data string) string {
	messageUUID := randomHex(8)

	attributes := MessageAttributes{
		WorkerName: workerName,
		Data:       data,
		Count:      0,
	}

	jsonAttributes, err := json.Marshal(attributes)
	if err != nil {
		log.Fatal(err)
	}

	err = e.Queue.Push(queue, string(jsonAttributes))
	if err != nil {
		log.Fatal(err)
	}

	// LOG DEBUG: fmt.Println("enqueuer", queue, string(jsonAttributes))

	return messageUUID
}
