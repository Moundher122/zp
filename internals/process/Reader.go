package process

import (
	"log"
	
	"github.com/cilium/ebpf"
)


func ReadWithoutRemoveFromMap(m *ebpf.Map, key uint32) (*Process, error) {
	var e Process

	// Lookup does NOT remove the entry
	if err := m.Lookup(&key, &e); err != nil {
		return nil, err
	}

	log.Printf("PID=%d PORT=%d\n", e.Pid, e.Port)
	return &e, nil
}
