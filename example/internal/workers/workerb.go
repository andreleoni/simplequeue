package workers

import (
	"fmt"
	"simplequeue"
)

type WorkerB struct {
	simplequeue.WorkerBase
}

func NewWorkerB() *WorkerB {
	workerB := WorkerB{}
	workerB.SetQueue("low")
	workerB.SetRegisterName("workerB")

	return &WorkerB{}
}

func (w WorkerB) Perform(data string) error {
	fmt.Println("processing worker B", data)

	return nil
}
