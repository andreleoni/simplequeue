package simplequeue

import (
	"fmt"
)

func Receiver(queueName string) []string {
	// DEBUG LOG: fmt.Println("looking for", queueName, "queued")

	data, err := RedisAdapter.Pop(queueName, 2)
	if err != nil && err.Error() != "redis: nil" {
		fmt.Println("error on messages Receiver: ", err)
	}

	return data
}
