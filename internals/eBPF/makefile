.PHONY: bind close clean vmlinux

CLANG = clang
LLC = llc
ARCH = $(shell uname -m | sed 's/x86_64/x86/' | sed 's/aarch64/arm64/')
LIBBPF_DIR = /usr/include
BPF_CFLAGS = -g -O2 -target bpf -D__TARGET_ARCH_$(ARCH)

bind:
	$(CLANG) $(BPF_CFLAGS) -I$(LIBBPF_DIR) -c bind/port.bind.bpf.c -o bind/port.bind.bpf.o

close:
	$(CLANG) $(BPF_CFLAGS) -I$(LIBBPF_DIR) -c close/port.close.bpf.c -o close/port.close.bpf.o

clean:
	rm -f bind/port.bind.bpf.o close/port.close.bpf.o

vmlinux:
	bpftool btf dump file /sys/kernel/btf/vmlinux format c > vmlinux.h
