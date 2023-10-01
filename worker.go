package simplequeue

type Worker interface {
	// Perform will process a job
	Perform(data string) error

	// PerformAsync will schedule a job on queue
	PerformAsync(data string) string

	// Retrieve worker options
	GetQueueName() string
	GetRegisterName() string

	// Set worker options
	SetQueueName(queueName string)
	SetRegisterName(registerName string)
}

type WorkerBase struct {
	registerName string
	queueName    string
}

func (w *WorkerBase) SetQueueName(queueName string) {
	w.queueName = queueName
}

func (w *WorkerBase) SetRegisterName(registerName string) {
	w.registerName = registerName
}

func (w *WorkerBase) GetQueueName() string {
	return w.queueName
}

func (w *WorkerBase) GetRegisterName() string {
	return w.registerName
}

func (w *WorkerBase) PerformAsync(data string) string {
	enqueuer := NewEnqueuer()

	return enqueuer.Enqueue(w.GetQueueName(), w.GetRegisterName(), data)
}
