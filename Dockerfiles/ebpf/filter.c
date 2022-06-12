#include <linux/bpf.h>
#include <linux/pkt_cls.h>
#include <bpf/bpf_helpers.h>

struct {
  __uint(type, BPF_MAP_TYPE_PERCPU_ARRAY);
  __uint(max_entries, 1);
  __type(key, int);
  __type(value, int);
} pkt_counter_ingress SEC(".maps");

struct {
  __uint(type, BPF_MAP_TYPE_PERCPU_ARRAY);
  __uint(max_entries, 1);
  __type(key, int);
  __type(value, int);
} pkt_counter_egress SEC(".maps");

struct {
  __uint(type, BPF_MAP_TYPE_PERF_EVENT_ARRAY);
  __uint(key_size, sizeof(int));
  __uint(value_size, sizeof(int));
} pkt_counter_events SEC(".maps");

static __inline int
count_packets(struct __sk_buff *skb, void *map)
{
  int key = 0;
  int *val = bpf_map_lookup_elem(map, &key);
  if (val == NULL) {
    return TC_ACT_SHOT;
  }
  *val = *val + 1;

  int msg = 0xefbeadde;
  bpf_perf_event_output(skb, &pkt_counter_events, BPF_F_CURRENT_CPU, &msg, sizeof(msg));

  return TC_ACT_OK;
}

SEC("tc-ingress") int
count_packets_ingress(struct __sk_buff *skb)
{
  return count_packets(skb, &pkt_counter_ingress);
}

SEC("tc-egress") int
count_packets_egress(struct __sk_buff *skb)
{
  return count_packets(skb, &pkt_counter_egress);
}

char __license[] SEC("license") = "GPL";
