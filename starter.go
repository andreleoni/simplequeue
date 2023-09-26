package simplequeue

import (
	"encoding/json"
	"fmt"
	"time"
)

func Starter(queuename string) error {
	defer func() error {
		if r := recover(); r != nil {
			fmt.Println("panic starting queue", r)
			return r.(error)
		}

		return nil
	}()

	for {
		receivedQueueData := Receiver(queuename)

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
