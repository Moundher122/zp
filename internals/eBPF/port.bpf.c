// port.bpf.c
#define __TARGET_ARCH_x86

#include "vmlinux.h"
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_core_read.h>
#include <bpf/bpf_tracing.h>
#include <bpf/bpf_endian.h>
#define AF_INET 2
struct event {
    __u32 pid;
    __u16 sport;
    __u16 dport;
};

struct {
    __uint(type, BPF_MAP_TYPE_RINGBUF);
    __uint(max_entries, 1 << 24);
} events SEC(".maps");

// Tracepoint for bind syscall
SEC("tracepoint/syscalls/sys_enter_bind")
int handle_bind(struct trace_event_raw_sys_enter *ctx)
{
    bpf_printk("=== BIND CALLED ===");
    
    struct event *e;
    __u32 pid;
    struct sockaddr_in addr;
    __u16 bind_port;
    
    pid = bpf_get_current_pid_tgid() >> 32;
    bpf_printk("PID: %u", pid);
    
    if (bpf_probe_read_user(&addr, sizeof(addr), (void *)ctx->args[1]) != 0) {
        bpf_printk("Failed to read sockaddr");
        return 0;
    }
    
    bpf_printk("Family: %u", addr.sin_family);
    
    if (addr.sin_family != AF_INET) {
        bpf_printk("Not AF_INET, skipping");
        return 0;
    }
    
    bind_port = bpf_ntohs(addr.sin_port);
    bpf_printk("Bind port: %u", bind_port);
    
    e = bpf_ringbuf_reserve(&events, sizeof(*e), 0);
    if (!e) {
        bpf_printk("Failed to reserve ringbuf");
        return 0;
    }
    
    e->pid   = pid;
    e->sport = bind_port;
    e->dport = 0;
    
    bpf_ringbuf_submit(e, 0);
    bpf_printk("Event submitted!");
    return 0;
}
char LICENSE[] SEC("license") = "GPL";