package simplequeue

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"os"
	"time"
)

func Starter(queuesName ...string) error {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	logger.Debug("usuário logado com sucesso")

	defer func() error {
		if r := recover(); r != nil {
			fmt.Println("panic starting queue", r)
			return r.(error)
		}

		return nil
	}()

	for {
		for _, queueName := range queuesName {
			// fmt.Println("looking for", queueName, "queued")

			receivedQueueData := Receiver(queueName)

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
