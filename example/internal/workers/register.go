package workers

import (
	"simplequeue"
)

func init() {
	simplequeue.Register(NewWorkerA())
	simplequeue.Register(NewWorkerB())
}
