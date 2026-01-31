package loadebpf

import (
	"github.com/cilium/ebpf/link"
)

func PinLink(link link.Link, path string) error {
	return link.Pin(path)
}
