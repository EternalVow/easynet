package _interface

type IEasyNet interface {
	OnStart(conn IConnection) error

	OnConnect(conn IConnection) error

	OnReceive(conn IConnection, ip IInputStream) ([]byte, error)

	OnShutdown(conn IConnection) error

	OnClose(conn IConnection, err error) error

	// todo to add more
}
