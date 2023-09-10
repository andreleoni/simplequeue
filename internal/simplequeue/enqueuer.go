package simplequeue

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
)

func Enqueuer(queue, workerName, data string) string {
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

	err = RedisAdapter.Push(queue, string(jsonAttributes))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("enqueuer", queue, string(jsonAttributes))

	return messageUUID
}

func randomHex(n int) string {
	bytes := make([]byte, n)

	return hex.EncodeToString(bytes)
}
