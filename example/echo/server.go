package main

import (
	"context"
	"fmt"
	"github.com/EternalVow/easynet"
	_interface "github.com/EternalVow/easynet/interface"
)

type Handler struct {
}

func (h Handler) OnStart(conn _interface.IConnection) error {
	fmt.Println("test OnStart")
	return nil
}

func (h Handler) OnConnect(conn _interface.IConnection) error {
	//netpollConn, ok:= conn.(netpoll.Connection)
	//if !ok {
	//	fmt.Println("test conn err")
	//}
	//
	//fmt.Println("test conn LocalAddr",netpollConn.LocalAddr())
	//fmt.Println("test conn RemoteAddr",netpollConn.RemoteAddr())

	return nil

}

func (h Handler) OnReceive(conn _interface.IConnection, stream _interface.IInputStream) ([]byte, error) {
	//netpollConn, ok:= conn.(netpoll.Connection)
	//if !ok {
	//	fmt.Println("test conn err")
	//}
	bytes := stream.Begin(nil)
	fmt.Println("test receive msg ", string(bytes))
	return []byte("1111111"), nil

}

func (h Handler) OnShutdown(conn _interface.IConnection) error {
	return nil
}

func (h Handler) OnClose(conn _interface.IConnection, err error) error {
	return nil
}

func main() {
	handler := &Handler{}
	easynet.NewEasyNetWithYamlConfig(context.Background(), "NetPoll", handler, "../../base/confg.yaml")
	//easynet.NewEasyNetWithYamlConfig(context.Background(), "Evio", handler, "../../base/confg.yaml")
	//easynet.NewEasyNetWithYamlConfig(context.Background(), "Net", handler, "../../base/confg.yaml")
	//easynet.NewEasyNetWithYamlConfig(context.Background(), "Gev", handler, "../../base/confg.yaml")
	//easynet.NewEasyNetWithYamlConfig(context.Background(), "Gnet", handler, "../../base/confg.yaml")

}
