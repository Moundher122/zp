package loadebpf

import (
	"log"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
)

func AtachToKernel(path string) (*ebpf.Collection, error) {
	spec, err := ebpf.LoadCollectionSpec("internals/eBPF/port.bpf.o")
	if err != nil {
		return nil, err
	}

	log.Println("=== programs ===")
	for name := range spec.Programs {
		log.Println(name)
	}

	log.Println("=== maps ===")
	for name := range spec.Maps {
		log.Println(name)
	}

	coll, err := ebpf.NewCollection(spec)
	if err != nil {
		return nil, err
	}
	log.Println(coll.Maps)
	prog := coll.Programs["handle_bind"]

	if prog == nil {
		log.Fatal("handle_tcp_connect not found")
	}

	lk, err := link.Tracepoint("syscalls", "sys_enter_bind", prog, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = PinLink(lk, path)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("eBPF program attached successfully")
	return coll, nil
}
