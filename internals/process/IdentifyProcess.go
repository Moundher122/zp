package process

import (
	"log"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/ringbuf"
)

type IdentifyProcess struct {
	Port int
	spec *ebpf.Map
}

func NewIdentifyProcess(port int, spec *ebpf.Map) *IdentifyProcess {
	return &IdentifyProcess{
		Port: port,
		spec: spec,
	}
}

func (ip *IdentifyProcess) Identify() *Process {
	rd, err := ringbuf.NewReader(ip.spec)
	if err != nil {
		log.Println("Error creating ring buffer reader:", err)
	}
	_, err = rd.Read()
	if err != nil {
		log.Println("Error reading from ring buffer:", err)
	}
	event, err := ReadWithoutRemoveFromMap(ip.spec, uint32(ip.Port))
	if err != nil {
		log.Println("Error parsing event data:", err)
	}
	return event
}
                                                                                                                                                    