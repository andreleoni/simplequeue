package workers

import (
	"fmt"
	"reflect"
	"simplequeue/internal/simplequeue"
)

func init() {
	workerA := WorkerA{}

	simplequeue.Register(workerA.StructName(), workerA)
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
	fmt.Println("PerformAsync", w.StructName(), w.QueueName(), data)

	return simplequeue.Enqueuer(w.QueueName(), w.StructName(), data)
}
