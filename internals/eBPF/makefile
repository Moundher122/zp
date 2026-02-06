CLANG = clang
LLC = llc
ARCH = $(shell uname -m | sed 's/x86_64/x86/' | sed 's/aarch64/arm64/')
LIBBPF_DIR = /usr/include
BPF_CFLAGS = -g -O2 -target bpf -D__TARGET_ARCH_$(ARCH)
all: port.bpf.o 
port.bpf.o: port.bpf.c
	$(CLANG) $(BPF_CFLAGS) -I$(LIBBPF_DIR) -c port.bpf.c -o port.bpf.o
clean:
	rm -f port.bpf.o 
vmlinux.h:
	bpftool btf dump file /sys/kernel/btf/vmlinux format c > vmlinux.h