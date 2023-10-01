package workers

import (
	"fmt"
	"simplequeue"
)

type WorkerA struct {
	simplequeue.WorkerBase
}

func NewWorkerA() *WorkerA {
	workerA := WorkerA{}
	workerA.SetQueueName("default")
	workerA.SetRegisterName("workerA")

	return &workerA
}

func (w *WorkerA) Perform(data string) error {
	fmt.Println("processing worker A", data)

	return nil
}
