package process

type Process struct {
	Pid  uint32
	Port uint16
}

func NewProcess(pid uint32, port uint16) *Process {
	return &Process{
		Pid:  pid,
		Port: port,
	}
}
