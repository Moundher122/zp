package loadebpf

import (
	"log"
	"os"

	"github.com/cilium/ebpf"
)

var (
	pinnedLinkPath = "/sys/fs/bpf/port_link"
	pinnedMapPath  = "/sys/fs/bpf/port_map"
)

func LoadEBPFProgram() (*ebpf.Map, error) {
	if _, err := os.Stat(pinnedLinkPath); os.IsNotExist(err) {
		coll, err := AtachToKernel(pinnedLinkPath, pinnedMapPath)
		return coll, err
	}
	coll, err := ebpf.LoadPinnedMap(pinnedMapPath, nil)
	if err != nil {
		log.Fatal("Failed to load pinned eBPF program:", err)
	}
	return coll, nil
}
