package workers

import "simplequeue/internal/simplequeue"

func init() {
	simplequeue.Register("workerA", WorkerA{})
}

type WorkerA struct {
	simplequeue.BaseWorker
}

func (w WorkerA) Perform(data string) error {
	return nil
}

func (w WorkerA) QueueName() string {
	return "high"
}

func (w WorkerA) Retry(count int) bool {
	return false
}
