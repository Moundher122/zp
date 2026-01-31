package loadebpf

import (
	"os"

	"github.com/cilium/ebpf"
)

var (
	pinnedLinkPath = "/sys/fs/bpf/port_link"
)

func LoadEBPFProgram() (*ebpf.Collection, error) {
	if _, err := os.Stat(pinnedLinkPath); os.IsNotExist(err) {
		coll, err := AtachToKernel(pinnedLinkPath)
		return coll, err
	}
	return nil, nil
}
