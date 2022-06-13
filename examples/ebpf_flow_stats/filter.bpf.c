#include <stdio.h>
#include <stdint.h>
#include <stdbool.h>
#include <byteswap.h>
#include <arpa/inet.h>
#include <linux/kernel.h>
#include <linux/if_ether.h>
#include <linux/in.h>
#include <linux/ip.h>
#include <linux/ipv6.h>
#include <linux/tcp.h>
#include <linux/udp.h>
#include <linux/bpf.h>
#include <linux/pkt_cls.h>
#include <bpf/bpf_helpers.h>

#define IP_MF     0x2000
#define IP_OFFSET 0x1FFF

#define assert_len(interest, end)                 \
  ({                                              \
    if ((unsigned long)(interest + 1) > data_end) \
      return TC_ACT_SHOT;                         \
  })

#define printk(fmt)                     \
  ({                                    \
    char msg[] = fmt;                   \
    bpf_trace_printk(msg, sizeof(msg)); \
  })

struct flowkey {
  // uint32_t ifindex;
  uint32_t daddr;
  uint16_t dport;
}  __attribute__ ((packed));

struct flowval {
  uint32_t cnt;
}  __attribute__ ((packed));

struct {
  __uint(type, BPF_MAP_TYPE_PERCPU_HASH);
  __uint(max_entries, 100);
  __type(key, struct flowkey);
  __type(value, struct flowval);
} flow_stats SEC(".maps");

struct {
  __uint(type, BPF_MAP_TYPE_PERF_EVENT_ARRAY);
  __uint(key_size, sizeof(int));
  __uint(value_size, sizeof(int));
} flow_events SEC(".maps");

static inline void record(const struct tcphdr *th, const struct iphdr *ih,
                          struct __sk_buff *skb)
{
  uint16_t dport = th->dest;
  uint32_t daddr = ih->daddr;
  struct flowval initval = {1};
  struct flowkey key = {0};
  key.daddr = daddr;
  key.dport = htons(dport);
  uint32_t *val = bpf_map_lookup_elem(&flow_stats, &key);
  if (val)
    *val = *val + 1;
  else
    bpf_map_update_elem(&flow_stats, &key, &initval, BPF_ANY);

  int msg = 0xefbeadde;
  bpf_perf_event_output(skb, &flow_events, BPF_F_CURRENT_CPU, &msg, sizeof(msg));
}

static inline int
process_ipv4_tcp(struct __sk_buff *skb)
{
  uint64_t data = skb->data;
  uint64_t data_end = skb->data_end;
  uint64_t pkt_len = 0;

  struct iphdr *ih = (struct iphdr *)(data + sizeof(struct ethhdr));
  assert_len(ih, data_end);
  pkt_len = data_end - data;

  uint8_t hdr_len = ih->ihl * 4;
  struct tcphdr *th = (struct tcphdr *)((char *)ih + hdr_len);
  assert_len(th, data_end);

  if (!(th->syn == 1 && th->ack == 0))
    return TC_ACT_SHOT;

  record(th, ih, skb);
  return TC_ACT_OK;
}

static inline int
process_ipv4_icmp(struct __sk_buff *skb)
{
  printk("icmp packet");
  return TC_ACT_OK;
}

static inline int
process_ipv4(struct __sk_buff *skb)
{
  uint64_t data = skb->data;
  uint64_t data_end = skb->data_end;
  uint64_t pkt_len = 0;

  struct iphdr *ih = (struct iphdr *)(data + sizeof(struct ethhdr));
  assert_len(ih, data_end);
  pkt_len = data_end - data;

  if (ih->ihl < 5)
    return TC_ACT_SHOT;

  switch (ih->protocol) {
  case IPPROTO_ICMP:
    return process_ipv4_icmp(skb);
  case IPPROTO_TCP:
    return process_ipv4_tcp(skb);
  default:
    return TC_ACT_OK;
  }
}

static inline int
process_ethernet(struct __sk_buff *skb)
{
  uint64_t data = skb->data;
  uint64_t data_end = skb->data_end;
  uint64_t pkt_len = 0;

  struct ethhdr *eth_hdr = (struct ethhdr *)data;
  assert_len(eth_hdr, data_end);
  pkt_len = data_end - data;

  switch (htons(eth_hdr->h_proto)) {
  case 0x0800:
    return process_ipv4(skb);
  default:
    return TC_ACT_SHOT;
  }
}

SEC("tc-ingress") int
count_packets_ingress(struct __sk_buff *skb)
{
  return process_ethernet(skb);
}

SEC("tc-egress") int
count_packets_egress(struct __sk_buff *skb)
{
  return process_ethernet(skb);
}

char __license[] SEC("license") = "GPL";
