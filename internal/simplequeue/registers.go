package simplequeue

var Registers = map[string]Worker{}

func Register(name string, worker Worker) {
	Registers[name] = worker
}
