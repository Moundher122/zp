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
	var e Process
	err = binary.Read(
		bytes.NewReader(record.RawSample),
		binary.LittleEndian,
		&e,
	)
	log.Printf("PID=%d PORT=%d\n", e.Pid, e.Port)
	return record.RawSample
}
