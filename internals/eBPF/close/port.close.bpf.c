#include "../vmlinux.h"
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_tracing.h>
#include <bpf/bpf_endian.h>

#define AF_INET 2

struct closePortEvent {
    __u32 pid;
    __u16 port;
};

struct {
    __uint(type, BPF_MAP_TYPE_RINGBUF);
    __uint(max_entries, 1 << 24);
} events SEC(".maps");

SEC("tracepoint/syscalls/sys_enter_close")
int handle_close(struct trace_event_raw_sys_enter *ctx)
{
    struct closePortEvent *e;
    __u32 pid;
    int fd;

    pid = bpf_get_current_pid_tgid() >> 32;
    fd  = (int)ctx->args[0];

    e = bpf_ringbuf_reserve(&events, sizeof(*e), 0);
    if (!e)
        return 0;

    e->pid  = pid;
    e->port = 0; 

    bpf_ringbuf_submit(e, 0);
    return 0;
}
