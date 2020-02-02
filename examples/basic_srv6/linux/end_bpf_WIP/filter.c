#include <stdint.h>
#include <linux/bpf.h>
#include <linux/if_ether.h>
#include <linux/if_packet.h>
#include <linux/ip.h>
#include "bpf_helpers.h"
#ifndef __section
#define __section(NAME) \
  __attribute__((section(NAME), used))
#endif

__section("prog")
int filter(struct __sk_buff *skb)
{
  /* void *data_end = (void *)(long)skb->data_end; */
	/* void *data = (void *)(long)skb->data; */
  /* if ((data + 0x0060 + 10) < data_end) { */
  /*   unsigned char* ptr = data; */
  /*   ptr[0x0060] = 'X'; */
  /* } */

  struct data {
    uint8_t r0;
    uint8_t r1;
    uint8_t r2;
    uint8_t r3;
  };
  struct data from;
  from.r0 = 1;

  /* int off = offsetof(struct iphdr, protocol); */
  int off = 40;
  int index = bpf_skb_store_bytes(skb, ETH_HLEN, &from, sizeof(from), 0);

  return BPF_OK;
  /* return BPF_OK; */
}
char __license[] __section("license") = "GPL";
