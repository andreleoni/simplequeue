package main

import (
	"encoding/json"
	"fmt"
	"log"
	"simplequeue/internal/simplequeue"

	_ "simplequeue/internal/workers"
)

func main() {
	fmt.Println(simplequeue.Registers)

	for i := 0; i < 100; i++ {
		for _, worker := range simplequeue.Registers {
			fmt.Println(worker.QueueName(), worker.PerformAsync("oi leoni"))
		}
	}

	receivedQueueData := simplequeue.Receiver("default")

	for {
		for _, queueData := range receivedQueueData {
			fmt.Println("count", len(queueData))
			messageData := simplequeue.MessageAttributes{}

			err := json.Unmarshal([]byte(queueData), &messageData)
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("attributes", messageData)

			gotStruct := simplequeue.Registers[messageData.WorkerName]

			gotStruct.Perform(messageData.Data)
		}

		receivedQueueData := simplequeue.Receiver("default")

		if len(receivedQueueData) < 10 {
			break
		}
	}

	fmt.Println("queues are empty")
}
