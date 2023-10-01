package simplequeue

import "fmt"

type Worker interface {
	Perform(data string) error
	PerformAsync(data string) string
	Queue() string
	RegisterName() string
	RetryCount() string
	SetQueue(queueName string)
	SetRegisterName(registerName string)
	SetRetryCount(retryCount int)
}

type WorkerBase struct {
	registerName string
	retryCount   int
	queue        string
}

func (w *WorkerBase) SetQueue(queueName string) {
	w.queue = queueName
}

func (w *WorkerBase) SetRegisterName(registerName string) {
	w.registerName = registerName
}

func (w *WorkerBase) SetRetryCount(retryCount int) {
	w.retryCount = retryCount
}

func (w *WorkerBase) Queue() string {
	return w.queue
}

func (w *WorkerBase) RegisterName() string {
	return w.registerName
}

func (w *WorkerBase) RetryCount() string {
	return w.registerName
}

func (w *WorkerBase) PerformAsync(data string) string {
	fmt.Println("PerformAsync", w.RegisterName(), w.Queue(), data)

	enqueuer := NewEnqueuer()

	return enqueuer.Enqueue(w.Queue(), w.RegisterName(), data)
}
