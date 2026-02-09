package loadebpf

import (
	"log"

	"github.com/cilium/ebpf"
	"github.com/cilium/ebpf/link"
)

func AtachToKernel(linkpath, mappath, ebpfpath string) (*ebpf.Map, error) {
	spec, err := ebpf.LoadCollectionSpec(ebpfpath)
	if err != nil {
		return nil, err
	}
	coll, err := ebpf.NewCollection(spec)
	if err != nil {
		return nil, err
	}
	prog := coll.Programs["handle_bind"]

	if prog == nil {
		log.Fatal("handle_tcp_connect not found")
	}

	lk, err := link.Tracepoint("syscalls", "sys_enter_bind", prog, nil)
	if err != nil {
		log.Fatal(err)
	}
	err = Pin(lk, linkpath)
	if err != nil {
		log.Fatal(err)
	}
	err = Pin(coll.Maps["events"], mappath)
	return coll.Maps["events"], nil
}
