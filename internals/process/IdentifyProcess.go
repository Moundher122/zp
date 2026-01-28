package process

import (
	"fmt"
	"os"
)

type IdentifyProcess struct {
	Port int
}

func NewIdentifyProcess(port int) *IdentifyProcess {
	return &IdentifyProcess{
		Port: port,
	}
}

func (p *IdentifyProcess) Identify() string {
	data, err := os.ReadFile("/proc/net/tcp")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	return "Process identified"
}
