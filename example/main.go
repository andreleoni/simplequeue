package main

import (
	"fmt"
	"reflect"
	"simplequeue"
	"time"
)

func main() {
	workerA := WorkerA{}
	simplequeue.Register(workerA.StructName(), workerA)

	workerB := WorkerB{}
	simplequeue.Register(workerB.StructName(), workerB)

	go func() {
		for {
			// Generate messages to test consumer
			for i := 0; i < 1; i++ {
				for _, worker := range simplequeue.Registers {
					go worker.PerformAsync(fmt.Sprint("my message ", time.Now()))
				}
			}

			time.Sleep(3 * time.Second)
		}
	}()

	err := simplequeue.Starter("default")

	fmt.Println("error got on server start", err)
}

type WorkerA struct {
}

func (w WorkerA) StructName() string {
	return fmt.Sprint(reflect.TypeOf(w))
}

func (w WorkerA) Perform(data string) error {
	fmt.Println("processing worker A", data)

	return nil
}

func (w WorkerA) QueueName() string {
	return "default"
}

func (w WorkerA) Retry(count int) bool {
	return false
}

func (w WorkerA) PerformAsync(data string) string {
	// fmt.Println("PerformAsync", w.StructName(), w.QueueName(), data)

	return simplequeue.Enqueuer(w.QueueName(), w.StructName(), data)
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
	// DEBUG LOG: fmt.Println("PerformAsync", w.StructName(), w.QueueName(), data)

	return simplequeue.Enqueuer(w.QueueName(), w.StructName(), data)
}
