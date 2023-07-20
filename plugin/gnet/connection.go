package gnet

import (
	"github.com/panjf2000/gnet/v2"
	"gnet"
)

type Connection struct {
	Conn gnet.Conn
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
