package main

import (
	"context"
	"log"
	"net"
	"syscall"
	"time"

	// Use miekg/dns patched with TPROXY support provided by Cilium.
	// Here they summarized the brief overview of what they did.
	// https://github.com/cilium/cilium/commit/4eb97b91843981b999f39e4acebf59681571f91f
	"github.com/cilium/dns"
	"golang.org/x/sys/unix"
)

func setupTransparentSocket(network, address string, c syscall.RawConn) error {
	var sysErr error

	err := c.Control(func(fd uintptr) {
		sysErr = unix.SetsockoptInt(int(fd), unix.SOL_IP, unix.IP_TRANSPARENT, 1)
		if sysErr != nil {
			return
		}

		sysErr = unix.SetsockoptInt(int(fd), unix.SOL_IP, unix.IP_RECVORIGDSTADDR, 1)
		if sysErr != nil {
			return
		}

		sysErr = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEADDR, 1)
		if sysErr != nil {
			return
		}

		sysErr = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1)
		if sysErr != nil {
			return
		}
	})

	if err != nil {
		return err
	}

	return sysErr
}

type DNSInterceptor struct {
}

func (di *DNSInterceptor) NewFQDNFirewall() {
}

func (di *DNSInterceptor) ServeDNS(w dns.ResponseWriter, req *dns.Msg) {
	src := w.RemoteAddr().String()
	dst := w.LocalAddr().String()
	proto := w.LocalAddr().Network()

	log.Printf("Request (from %s to %s proto %s) ===\n%s\n", src, dst, proto, req.String())

	var client *dns.Client
	switch proto {
	case "udp":
		client = &dns.Client{Net: "udp", SingleInflight: false}
	case "tcp":
		client = &dns.Client{Net: "tcp", SingleInflight: false}
	default:
		panic("got DNS query from unknown network")
	}

	// preserve original ID for later response and use another ID for outgoing request
	origId := req.Id
	req.Id = dns.Id()

	res, _, err := client.Exchange(req, dst)
	if err != nil {
		log.Println("Failed to forward DNS query")
		return
	}

	log.Printf("Response (from %s to %s proto %s) ===\n%s\n", dst, src, proto, res.String())

	res.Id = origId
	w.WriteMsg(res)
}

func main() {
	listenConf := &net.ListenConfig{
		Control: setupTransparentSocket,
	}

	tcpListener, err := listenConf.Listen(context.Background(), "tcp4", ":53")
	if err != nil {
		panic(err)
	}

	udpListener, err := listenConf.ListenPacket(context.Background(), "udp4", ":53")
	if err != nil {
		panic(err)
	}

	di := &DNSInterceptor{}

	tcpServer := &dns.Server{
		Net:      "tcp4",
		Listener: tcpListener,
		Handler:  di,
	}

	udpServer := &dns.Server{
		Net:        "udp4",
		PacketConn: udpListener,
		Handler:    di,
		SessionUDPFactory: &SessionUDPFactory{ipv4Enabled: true, ipv6Enabled: false},
	}

	go tcpServer.ActivateAndServe()
	go udpServer.ActivateAndServe()

	log.Println("Listening...")

	for {
		time.Sleep(1 * time.Second)
	}
}
