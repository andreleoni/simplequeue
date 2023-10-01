package simplequeue

import "fmt"

type Worker interface {
	Perform(data string) error
	PerformAsync(data string) string
	Queue() string
	StructName() string
	SetOptions(opts WorkerOptions)
}

type WorkerOptions struct {
	retryCount int
	queue      string
	structName string
}

func NewWorkerOptions(retry int, queue, structName string) WorkerOptions {
	return WorkerOptions{
		retryCount: retry,
		queue:      queue,
		structName: structName,
	}
}

type WorkerBase struct {
	Opts WorkerOptions
}

func (w *WorkerBase) SetOptions(opts WorkerOptions) {
	w.Opts = opts
}

func (w *WorkerBase) Queue() string {
	return w.Opts.queue
}

func (w *WorkerBase) StructName() string {
	return w.Opts.structName
}

func (w *WorkerBase) PerformAsync(data string) string {
	fmt.Println("PerformAsync", w.StructName(), w.Queue(), data)

	enqueuer := NewEnqueuer()

	return enqueuer.Enqueue(w.Queue(), w.StructName(), data)
}
