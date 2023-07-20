package evio

import "github.com/tidwall/evio"

type Connection struct {
	Conn evio.Conn
}

func (c Connection) RemoteAddr() string {
	return c.Conn.RemoteAddr().String()
}

func (c Connection) Send(bs []byte) (n int, err error) {
	return 0, err
}

func (c Connection) Close() (err error) {
	return nil
}
