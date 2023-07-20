package _interface

type IConnection interface {
	RemoteAddr() string
	Send(bs []byte) (n int, err error)
	Close() (err error)
}
