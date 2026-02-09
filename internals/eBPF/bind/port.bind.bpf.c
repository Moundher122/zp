#include "../vmlinux.h"
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_tracing.h>
#include <bpf/bpf_endian.h>

#define AF_INET 2

struct OpenPortEvent {
    __u32 pid;
    __u16 port;
};

struct {
    __uint(type, BPF_MAP_TYPE_RINGBUF);
    __uint(max_entries, 1 << 24);
} events SEC(".maps");

SEC("tracepoint/syscalls/sys_enter_bind")
int handle_bind(struct trace_event_raw_sys_enter *ctx)
{
    struct sockaddr sa = {};
    struct sockaddr_in *sin;
    struct OpenPortEvent *e;
    __u32 pid;

    pid = bpf_get_current_pid_tgid() >> 32;
    if (bpf_probe_read_user(&sa, sizeof(sa), (void *)ctx->args[1])) {
        return 0;
    }
    if (sa.sa_family != AF_INET) {
        return 0;
    }
    struct sockaddr_in sin4 = {};
    if (bpf_probe_read_user(&sin4, sizeof(sin4), (void *)ctx->args[1])) {
        return 0;
    }
    e = bpf_ringbuf_reserve(&events, sizeof(*e), 0);
    if (!e)
        return 0;
    e->pid  = pid;
    e->port = bpf_ntohs(sin4.sin_port);
    bpf_ringbuf_submit(e, 0);
    return 0;
}

char LICENSE[] SEC("license") = "GPL";
