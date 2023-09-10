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

	for _, worker := range simplequeue.Registers {
		fmt.Println(worker.QueueName(), worker.PerformAsync("oi leoni"))
	}

	receivedQueueData := simplequeue.Receiver("default")

	for _, queueData := range receivedQueueData {
		messageData := simplequeue.MessageAttributes{}

		err := json.Unmarshal([]byte(queueData), &messageData)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println("attributes", messageData)

		gotStruct := simplequeue.Registers[messageData.WorkerName]

		gotStruct.Perform(messageData.Data)
	}
}
