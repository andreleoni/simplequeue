package workers

import "simplequeue/internal/simplequeue"

func init() {
	simplequeue.Register("workerB", WorkerB{})
}

type WorkerB struct {
	simplequeue.BaseWorker
}

func (w WorkerB) Perform(data string) error {
	return nil
}

func (w WorkerB) QueueName() string {
	return "low"
}

func (w WorkerB) Retry(count int) bool {
	return true
}
