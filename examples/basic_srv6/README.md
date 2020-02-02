
# Segment Routing IPv6 Examples

- **vpn-v4-per-ce**: IPv4 L3VPN [T.Encaps, End, End.DX4]
- **vpn-v6-per-ce**: IPv6 L3VPN [T.Encaps, End, End.DX6]
- **vpn-v4-per-vrf**: IPv4 L3VPN [T.Encaps, End, End.DT4]  (WIP)
- **vpn-v6-per-vrf**: IPv6 L3VPN [T.Encaps, End, End.DT6]
- **l2vpn**: L2VPN [T.Encaps.L2, End, End.DX4]   (XXX)
- **transit**: Transit Function Test with FRR, [T.Encaps, T.Insert, End, End.X]
- **vrf-redirect**: End.T Evaluation linux only without FRR, [End.T]  (XXX)
- **binding-sid**: End.B6,End.B6.Encaps Evaluation. [End.B6, End.B6.Encaps]  (WIP End.B6.Insert)
<!-- - **sfc**: Service Chaining with End.AM -->

## Functions I want to test on linux
- [x] T.Insert: I don't understand how-to-use
- [x] T.Encaps
- [x] T.Encaps.L2: I don't understand how-to-use
- [x] End
- [x] End.X
- [ ] End.T: linux kernel can't perform End.T fine. (XXX)
- [x] End.DX6
- [x] End.DX4
- [x] End.DX2
- [x] End.DT6
- [ ] End.DT4: linux kernel isn't support yet
- [x] End.B6: linux kernel isn't working fine..? or I don't understand how-to-use.
- [x] End.B6.Encaps
- [ ] End.AM: linux kernel isn't support yet. **srext only** masquerading proxy
- [ ] End.AD: linux kernel isn't support yes. **srext only** dynamic proxy
- [ ] End.AS2: linux kernel isn't support yes. static proxy
- [ ] End.AS4: linux kernel isn't support yes. static proxy
- [ ] End.AS6: linux kernel isn't support yes. static proxy

## Funcに関しての説明
- **T.Encaps:**<br>
  この操作はトランジットノード
	(すなわち、パケットを通してSRv6をサポートするルーターですが、
	ノード自体はSegmentリストにありません)で実行され、
	IPv6ヘッダーとSRHヘッダーをパケットの外側の層に追加し、
	新しいものを定義できます。セグメントのリスト。
	パケットは、新しいIPv6ヘッダーのSRHに従って最初に転送されます。
- **End:**<br>
  この操作では、Segment Leftを0（最後のホップではない）
	にする必要があります。これにより、Segment Leftが1減少し、
	IPv6パケットの宛先アドレスが最も一般的なSRv6操作である次の
	Segmentに更新されます。SR MPLSのプレフィックスSIDと同等です。
- **End.X:**<br>
  この操作は基本的にEnd操作と同じですが、
  処理したパケットを指定したネクストホップアドレスに
	送信できる点が異なります。
	SR MPLSのAdj-SIDと同じです。
- **End.DX4:**<br>
  この操作ではSegment Leftが0であり、
	パケットがIPv4パケットをカプセル化しているため、
	外側のIPv6ヘッダーは削除され、内部のIPv4パケットは
	指定されたネクストホップアドレスに転送されます。
	VPNv4のCEごとのラベルに相当します。
- **End.DX6:**<br>
  この操作では、Segment Leftを0にしてIPv6パケットを
	パケットにカプセル化し、外側のIPv6ヘッダーを削除して、
	内部IPv6パケットを指定されたネクストホップアドレスに転送します。
	VPNv6 Per-CEラベルと同等です。
- **End.B6:**<br>
  既存のSRHに基づいて新しいSRHを挿入し、
	新しいSegment listを定義すると、挿入された新しいSRHに
	従ってデータパケットが最初に転送されます。Binding-SIDと同等です。
- **End.B6.Encaps:**<br>
  この操作は基本的にEnd.B6と同じですが、違いは、
	この操作では単にSRHルーティングヘッダーを追加するのではなく、
	新しいIPv6ヘッダーとSRHをパケットの外側のレイヤーに追加することです。

## 参照

**Funcの概要に関してはこれらの画像がわかりやすい.(thx kamataさん)**<br>

| Function | Location | Description | Works-for |
| :------- | :------- | :---------- | :-------- |
| End           | Core | DestとSRHを書き換えて, Next-hopをRIBから探して送る                    | Prefix-SID            |
| End.X         | Core | DestとSRHを書き換えて, 決められたNextHopへ送る                        | Adjacency-SID         |
| End.T         | Core | DestとSRHを書き換えて, NextHopを指定したRIBから探して送る             | Multi-table Operation |
| End.DX2       | Edge | SRHを外して, 決められた送信IFへ送る (NH=59)                           | L2VPN                 |
| End.DX6       | Edge | SRHを外して, 決められたIPv6 NextHopへ送る (NH=41)                     | VPNv6 Per-CE Label    |
| End.DX4       | Edge | SRHを外して, 決められたIPv4 NextHopへ送る (NH=4)                      | VPNv4 Per-CE Label    |
| End.DT6       | Edge | SRHを外して, IPv6 NextHopを指定したRIBから探して送る (NH=41)          | VPNv6 Per-VRF Label   |
| End.DT4       | Edge | SRHを外して, IPv4 NextHopを指定したRIBから探して送る (NH=4)           | VPNv4 Per-VRF Label   |
| End.B6        | Edge | SRHは触らず, 新しいSID List(SRH)を挿入して, その先頭に送る            | Binding SID           |
| End.B6.Encaps | Edge | SRHを書き換えて, 新しいSID List(OuterHdr)でEncapして, その先頭に送る  | Binding SID (Encap)   |
| End.BM        | Edge | DestとSRHを書き換えて, Labelを付与して, その先頭に送る                | SRv6/SR-MPLS Binding  |
| End.S         | Core | 一番最後(or複数)のSIDでTable検索し, NextHopを探して送る               | ICN                   |
| End.AS        | Core | OuterHdrを外して, 決められた送信IFへ送る. <br>決められた受信IFに入ってきたPktにOuterHdrを付与し, その先頭へ送る| Service-Chaining (Proxy)  |
| End.AM        | Core | DestとSRHを書き換えて, 決められた送信IFへ送る. <br>決められた受信IFに入ってきたPktにSRHを付与し, その先頭に送る. | Service-Chaining (Masq) |
| T             | Core | 通常のIPv6 Routing                                    | - |
| T.Insert      | Core | 新しいSRHを挿入して, その先頭に送る                   | - |
| T.Encaps      | Core | 新しいIPv6 Hdr(SRHつき)を挿入して, その先頭に送る(L3)􏲟| - |
| T.Encaps.L2   | Core | 新しいIPv6 Hdr(SRHつき)を挿入して, その先頭に送る(L2)􏲟| - |

**iproute2 CLIに関してはこれらがわかりやすい.(thx ebikenさん)**<br>
```
# ip -6 route add <segment> encap seg6local action <action> <params> (dev <device> | via <nexthop>) [table localsid]
# ip -6 route add fc00::1/128 encap seg6local \
			action End                                    via 2001:db8::1
			action End.X         nh6 fc00::1:1            via 2001:db8::1
			action End.T         table 100                via 2001:db8::1
			action End.DX2       oif lxcbr0               via 2001:db8::1
			action End.DX4       nh4 10.0.3.254           via 2001:db8::1
			action End.DX6       nh6 fc00::1:1            via 2001:db8::1
			action End.DT6       table 100                via 2001:db8::1
			action End.B6        srh segs beaf::1,beaf::2 via 2001:db8::1
			action End.B6.Encaps srh segs beaf::1,beaf::2 via 2001:db8::1

# ip -6 route add <prefix> encap seg6 mode <encapmode> segs <segments> [hmac <keyid>] (dev <device> | via <nexthop>)
# ip -6 route add fc00:b::10/128 encap seg6 mode inline segs fc00:3::11,fc00:3::12,fc00:3::13 via fc00:a::a
# ip -6 route add fc00:b::10/128 encap seg6 mode encap  segs fc00:3::11,fc00:3::12,fc00:3::13 via fc00:a::a

# ip sr tunsrc set <addr>
# ip sr hmac set <keyid> <algorithm>
```

- https://www.sdnlab.com/22842.html
- https://www.janog.gr.jp/meeting/janog40/application/files/2415/0051/7614/janog40-sr-kamata-takeda-00.pdf
- https://www.slideshare.net/kentaroebisawa/zebra-srv6-cli-on-linux-dataplane-enog49
- http://www.segment-routing.net/open-software/linux/
- http://www.segment-routing.net/open-software/vpp/

