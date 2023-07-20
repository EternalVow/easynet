package net

import (
	"net"
)

type Connection struct {
	Conn net.Conn
}

func (c Connection) RemoteAddr() string {
	return c.Conn.RemoteAddr().String()
}

func (c Connection) Send(bs []byte) (n int, err error) {
	return c.Conn.Write(bs)
}

func (c Connection) Close() (err error) {
	return c.Conn.Close()
}
