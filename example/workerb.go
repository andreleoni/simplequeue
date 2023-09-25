package main

import (
	"fmt"
	"reflect"
	"simplequeue"
)

func init() {
	workerB := WorkerB{}

	simplequeue.Register(workerB.StructName(), workerB)
}

type WorkerB struct {
}

func (w WorkerB) StructName() string {
	return fmt.Sprint(reflect.TypeOf(w))
}

func (w WorkerB) Perform(data string) error {
	fmt.Println("processing worker B", data)

	return nil
}

func (w WorkerB) QueueName() string {
	return "default"
}

func (w WorkerB) Retry(count int) bool {
	return true
}

func (w WorkerB) PerformAsync(data string) string {
	fmt.Println("PerformAsync", w.StructName(), w.QueueName(), data)

	return simplequeue.Enqueuer(w.QueueName(), w.StructName(), data)
}
