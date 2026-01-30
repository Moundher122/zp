package main

import (
	"zp/internals/loadebpf"
	"zp/internals/process"

	"log"

	"github.com/cilium/ebpf/rlimit"
)

func main() {
	if err := rlimit.RemoveMemlock(); err != nil {
		log.Fatal(err)
	}
	ch := make(chan interface{})
	go check(ch)
	m := <-ch
	log.Println("Identified eBPF Map Spec:", m)
}
func check(ch chan interface{}) {
	spec, err := loadebpf.LoadEBPFProgram()
	if err != nil {
		println("Error loading eBPF program:", err.Error())
		return
	}
	proc := process.NewIdentifyProcess(3000, spec)
	for {
		result := proc.Identify()
		if result != nil {
		}
	}
}
