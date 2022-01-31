package main

import (
	"context"
	"log"
	"net"
	"syscall"
	"net/http"
	"net/http/httputil"

	"golang.org/x/sys/unix"
)

func setupTransparentSocket(network, address string, c syscall.RawConn) error {
	var sysErr error

	err := c.Control(func(fd uintptr) {
		sysErr = unix.SetsockoptInt(int(fd), unix.SOL_IP, unix.IP_TRANSPARENT, 1)
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

func dialContextWithTransparentSocket(ctx context.Context, network, address string) (net.Conn, error) {
	return (&net.Dialer{
		Control: setupTransparentSocket,
	}).DialContext(ctx, network, address)
}

type HTTPInterceptor struct {
}

func (hi *HTTPInterceptor) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	src := req.RemoteAddr
	dst := req.Context().Value(http.LocalAddrContextKey)

	reqBin, err := httputil.DumpRequest(req, true)
	if err != nil {
		log.Printf("Failed to dump http request: %s\n", err)
		return
	}

	log.Printf("Request (from %s to %s) ===\n%s\n", src, dst, string(reqBin))

	// We are only listerning on the http (80), so this works
	req.RequestURI = ""
	req.URL.Scheme = "http"
	req.URL.Host = req.Host

	tr := &http.Transport{
		DialContext: dialContextWithTransparentSocket,
	}

	client := &http.Client{Transport: tr}

	res, err := client.Do(req)
	if err != nil {
		log.Printf("Failed to forward HTTP request (from %s to %s): %s\n", src, dst, err)
		return
	}

	defer res.Body.Close()

	resBin, err := httputil.DumpResponse(res, true)
	if err != nil {
		log.Printf("Failed to dump http response: %s\n", err)
		return
	}

	log.Printf("Response (from %s to %s) ===\n%s\n", dst, src, string(resBin))

	res.Write(w)
}

func main() {
	listenConf := &net.ListenConfig{
		Control: setupTransparentSocket,
	}

	listener, err := listenConf.Listen(context.Background(), "tcp4", ":80")
	if err != nil {
		log.Fatal(err)
	}

	server := &http.Server{
		Handler: &HTTPInterceptor{},
	}

	server.Serve(listener)
}
