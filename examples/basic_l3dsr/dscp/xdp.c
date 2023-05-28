#include <linux/bpf.h>
#include <linux/if_ether.h>
#include <linux/icmp.h>
#include <linux/ip.h>
#include <linux/in.h>
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_endian.h>

#ifndef DSCP
#define DSCP 10
#endif

static __always_inline __u16 csum_fold_helper(__u64 csum) {
  int i;
#pragma unroll
  for (i = 0; i < 4; i ++) {
    if (csum >> 16)
      csum = (csum & 0xffff) + (csum >> 16);
  }
  return ~csum;
}

static __always_inline void ipv4_csum(void *data_start, int data_size,  __u64 *csum) {
  *csum = bpf_csum_diff(0, 0, data_start, data_size, *csum);
  *csum = csum_fold_helper(*csum);
}

static __always_inline int process_ipv4(struct xdp_md* ctx,
  __u64 data, __u64 data_end) {
  __u8 src_mac[ETH_ALEN];
  struct ethhdr *eth;
  struct iphdr *iph;

  __u32 dst_addr = 0xa000201;
  __u64 csum = 0;

  eth = (struct ethhdr*)(data);

  if ((__u64)(eth + 1) > data_end)
    return XDP_DROP;

  iph = (struct iphdr*)(data + sizeof(struct ethhdr));

  if ((__u64)(iph + 1) > data_end)
    return XDP_DROP;

  if (iph->daddr == bpf_htonl(0x8e000001)) {
    iph->tos = DSCP;
    iph->daddr = bpf_htonl(dst_addr);
    iph->check = 0;
    ipv4_csum(iph, sizeof(struct iphdr), &csum);
    iph->check = csum;

    __builtin_memcpy(src_mac, eth->h_source, ETH_ALEN);
    __builtin_memcpy(eth->h_source, eth->h_dest, ETH_ALEN);
    __builtin_memcpy(eth->h_dest, src_mac, ETH_ALEN);
  
    return XDP_TX;
  }
  return XDP_PASS;
}

static __always_inline int process_eth(struct xdp_md* ctx) {
  __u64 data = ctx->data;
  __u64 data_end = ctx->data_end;
  struct ethhdr *eth;

  eth = (struct ethhdr*)data;

  if ((__u64)(eth + 1) > data_end)
    return XDP_DROP;

  if (eth->h_proto == bpf_htons(ETH_P_IP)) {
    return process_ipv4(ctx, data, data_end);
  }
  return XDP_PASS;
}

SEC("xdp-lb")
int entry(struct xdp_md *ctx) {
  int ret = process_eth(ctx);
  return ret;
}

char __license[] SEC("license") = "GPL";
