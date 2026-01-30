package process

type Process struct {
	Pid   int
	sport int
	dport int
}

func NewProcess(pid, sport, dport int) *Process {
	return &Process{
		Pid:   pid,
		sport: sport,
		dport: dport,
	}
}
