package main

import (
	"context"
	"fmt"
	"github.com/EternalVow/easynet"
	_interface "github.com/EternalVow/easynet/interface"
)

type Handler struct {
}

func (h Handler) OnStart(conn interface{}) error {
	fmt.Println("test OnStart")
	return nil
}

func (h Handler) OnConnect(conn interface{}) error {
	//netpollConn, ok:= conn.(netpoll.Connection)
	//if !ok {
	//	fmt.Println("test conn err")
	//}
	//
	//fmt.Println("test conn LocalAddr",netpollConn.LocalAddr())
	//fmt.Println("test conn RemoteAddr",netpollConn.RemoteAddr())

	return nil

}

func (h Handler) OnReceive(conn interface{}, stream _interface.IInputStream) ([]byte, error) {
	//netpollConn, ok:= conn.(netpoll.Connection)
	//if !ok {
	//	fmt.Println("test conn err")
	//}
	bytes := stream.Begin(nil)
	fmt.Println("test receive msg ", string(bytes))
	return []byte("1111111"), nil

}

func (h Handler) OnShutdown(conn interface{}) error {
	return nil
}

func (h Handler) OnClose(conn interface{}, err error) error {
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
