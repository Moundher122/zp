package process

import (
	"bytes"
	"encoding/binary"
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

func (ip *IdentifyProcess) Identify() []byte {
	rd, err := ringbuf.NewReader(ip.spec)
	if err != nil {
		log.Println("Error creating ring buffer reader:", err)
	}
	record, err := rd.Read()
	if err != nil {
		log.Println("Error reading from ring buffer:", err)
	}
	var event Process
	err = binary.Read(bytes.NewBuffer(record.RawSample), binary.LittleEndian, &event)
	if err != nil {
		log.Println("Error parsing event data:", err)
	}
	return record.RawSample
}
