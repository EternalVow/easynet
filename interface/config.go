package _interface

type IConfig interface {
	GetProtocol() string
	GetIp() string
	GetPort() int32
}
