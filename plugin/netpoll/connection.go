package netpoll

import (
	"github.com/cloudwego/netpoll"
)

type Connection struct {
	Conn netpoll.Connection
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
