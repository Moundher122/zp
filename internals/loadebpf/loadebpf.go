package loadebpf

import (
	"log"
	"os"

	"github.com/cilium/ebpf"
)

func LoadEBPFProgram(pinnedLinkPath, pinnedMapPath string) (*ebpf.Map, error) {
	if _, err := os.Stat(pinnedLinkPath); os.IsNotExist(err) {
		RBMAP, err := AtachToKernel(pinnedLinkPath, pinnedMapPath, "internals/eBPF/bind/port.bind.bpf.o")
		return RBMAP, err
	}
	RBMAP, err := ebpf.LoadPinnedMap(pinnedMapPath, nil)
	if err != nil {
		log.Fatal("Failed to load pinned eBPF program:", err)
	}
	return RBMAP, nil
}
