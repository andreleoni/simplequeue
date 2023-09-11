package simplequeue

import (
	"fmt"
	"log"
)

func Receiver(queueName string) []string {
	fmt.Println("looking for", queueName, "queued")

	data, err := RedisAdapter.Pop(queueName, 2)
	if err != nil {
		log.Fatal(err)
	}

	return data
}
