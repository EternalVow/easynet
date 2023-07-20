package easynet

import (
	"context"
	"fmt"
	_interface "github.com/EternalVow/easynet/interface"
	"testing"
)

type Handler struct {
}

func (h Handler) OnStart(conn _interface.IConnection) error {
	return nil
}

func (h Handler) OnConnect(conn _interface.IConnection) error {
	return nil

}

func (h Handler) OnReceive(conn _interface.IConnection, stream _interface.IInputStream) ([]byte, error) {
	return nil, nil

}

func (h Handler) OnShutdown(conn _interface.IConnection) error {
	return nil
}

func (h Handler) OnClose(conn _interface.IConnection, err error) error {
	return nil
}

func TestEasyNet(t *testing.T) {
	config := NewDefaultNetConfig("tcp", "127.0.0.1", 9011)
	handler := &Handler{}
	net := NewEasyNet(context.Background(), "NetPoll", config, handler)
	fmt.Println(net)
}

func TestEasyNetWithYamlConfig(t *testing.T) {
	handler := &Handler{}
	net := NewEasyNetWithYamlConfig(context.Background(), "Evio", handler, "./base/confg.yaml")
	fmt.Println(net)
}
