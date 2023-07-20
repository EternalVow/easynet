package gev

import (
	"github.com/Allenxuxu/gev"
)

type Connection struct {
	Conn *gev.Connection
}

func (c Connection) RemoteAddr() string {
	return c.Conn.PeerAddr()
}

func (c Connection) Send(bs []byte) (n int, err error) {
	err = c.Conn.Send(bs)
	if err != nil {
		return 0, err
	}
	return len(bs), nil
}

func (c Connection) Close() (err error) {
	return c.Conn.Close()
}
