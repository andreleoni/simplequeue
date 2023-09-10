package simplequeue

type Worker interface {
	Perform(data string) error
	QueueName() string
	Retry(count int) bool
	PerformAsync(data string) string
	StructName() string
}
