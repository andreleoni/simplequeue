package main

import (
	"fmt"
	"simplequeue/internal/simplequeue"

	_ "simplequeue/internal/workers"
)

func main() {
	fmt.Println(simplequeue.Registers)

	for _, worker := range simplequeue.Registers {
		fmt.Println(worker.QueueName(), worker.PerformAsync("oi leoni"))
	}

	fmt.Println(simplequeue.Receiver("default"))
	fmt.Println(simplequeue.Receiver("low"))
	fmt.Println(simplequeue.Receiver("high"))

}
