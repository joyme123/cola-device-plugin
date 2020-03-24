package server

import (
	"net"
	"time"
)

//UnixDial dials to a unix socket using net.DialTimeout
func UnixDial(addr string, timeout time.Duration) (net.Conn, error) {
	return net.DialTimeout("unix", addr, timeout)
}
