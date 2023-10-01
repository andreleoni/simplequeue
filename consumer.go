package simplequeue

import (
	"encoding/json"
	"fmt"
	"simplequeue/queue"
	"time"
)

type Consumer struct {
	Queue queue.Queue
	Opts  ConsumerOpts
}

type ConsumerOpts struct {
	QueuesName []string
}

func NewConsumer(queueService queue.Queue, opts ConsumerOpts) *Consumer {
	consumer := Consumer{Queue: queueService, Opts: opts}

	return &consumer
}

func (c *Consumer) Consume() error {
	// logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
	// 	Level: slog.LevelInfo,
	// }))

	// logger.Debug("usuário logado com sucesso")

	defer func() error {
		if r := recover(); r != nil {
			fmt.Println("panic starting queue", r)
			return r.(error)
		}

		return nil
	}()

	for {
		for _, queueName := range c.Opts.QueuesName {
			fmt.Println("looking for", queueName, "queued")

			receivedQueueData, err := c.Queue.Pop(queueName, 2)

			if err != nil && err.Error() != "redis: nil" {
				fmt.Println("error on messages Receiver: ", err)
				continue
			}

			for _, queueData := range receivedQueueData {
				messageData := MessageAttributes{}

				err := json.Unmarshal([]byte(queueData), &messageData)
				if err != nil {
					fmt.Println("error unmarshalling queue data", err)
					continue
				}

				gotStruct := Registers[messageData.WorkerName]
				if gotStruct == nil {
					fmt.Println("worker not found", messageData.WorkerName)
					continue
				}

				go gotStruct.Perform(messageData.Data)
			}

			time.Sleep(500 * time.Millisecond)
		}
	}
}
