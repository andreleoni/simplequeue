package simplequeue

var Registers = map[string]Worker{}

func Register(worker Worker) {
	Registers[worker.GetRegisterName()] = worker
}
