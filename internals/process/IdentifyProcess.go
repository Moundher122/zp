package process

import (
	"log"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/ringbuf"
	"github.com/dgraph-io/badger/v4"
)

type IdentifyProcess struct {
	spec *ebpf.Map
	db   *badger.DB
}

func NewIdentifyProcess(spec *ebpf.Map, db *badger.DB) *IdentifyProcess {
	return &IdentifyProcess{
		spec: spec,
		db:   db,
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
	event, err := ReadFromRingBuf(rd, ip.db)
	if err != nil {
		log.Println("Error parsing event data:", err)
	}
	return event
}
