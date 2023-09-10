package simplequeue

import (
	"fmt"
	"reflect"
)

type Worker interface {
	Perform(data string) error
	QueueName() string
	Retry(count int) bool
	PerformAsync(data string) string
	StructName() string
}

type BaseWorker struct{}

func (w BaseWorker) QueueName() string {
	return "default"
}

func (w BaseWorker) StructName() string {
	return fmt.Sprint(reflect.TypeOf(w))
}

func (w BaseWorker) PerformAsync(data string) string {
	fmt.Println(w.StructName(), w.QueueName(), data)

	return Enqueuer(w.QueueName(), w.StructName(), data)
}
