// Package proxy provides support for a variety of protocols to proxy network
// data.
package client

import (
	"github.com/litecy/goproxy/pkg/core/lib/socks5"
	socks52 "github.com/litecy/goproxy/pkg/core/proxy/client/socks5"
	"net"
	"time"
)

// A Dialer is a means to establish a connection.
type Dialer interface {
	// Dial connects to the given address via the proxy.
	DialConn(conn *net.Conn, network, addr string) (err error)
}

// Auth contains authentication parameters that specific Dialers may require.
type Auth struct {
	User, Password string
}

func SOCKS5(timeout time.Duration, auth *Auth) (Dialer, error) {
	var a *socks5.UsernamePassword
	if auth != nil {
		a = &socks5.UsernamePassword{auth.User, auth.Password}
	}
	d := socks52.NewDialer(a, timeout)
	return d, nil
}
