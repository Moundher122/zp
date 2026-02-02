package main

import (
	"log"
	"zp/cmd"

	"github.com/cilium/ebpf/rlimit"
)

func main() {
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatal(err)
	}
	cmd.Execute()
}
