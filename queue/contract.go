package queue

type Queue interface {
	Push(queueName, data string) error
	Pop(queueName string, count int) ([]string, error)
}
