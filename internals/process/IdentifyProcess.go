package process

import (
	"log"
	converter "zp/pkg/Converter"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/ringbuf"
)

type IdentifyProcess struct {
	Port int
	spec *ebpf.Collection
}

func NewIdentifyProcess(port int, spec *ebpf.Collection) *IdentifyProcess {
	return &IdentifyProcess{
		Port: port,
		spec: spec,
	}
}

func (ip *IdentifyProcess) Identify() *ebpf.Map {
	the := ip.spec.Maps["events"]
	log.Println("Identifying process for port:", converter.PortConverter(6379))
	log.Println(the)
	rd, err := ringbuf.NewReader(the)
	if err != nil {
		log.Println("Error creating ring buffer reader:", err)
	}
	record, err := rd.Read()
	if err != nil {
		log.Println("Error reading from ring buffer:", err)
	}
	log.Println(record.RawSample)
	log.Println(the)
	return the
}
