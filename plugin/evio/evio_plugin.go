package evio

import (
	"context"
	"net"

	"github.com/baickl/logger"
	"github.com/EternalVow/easynet/interface"
)

type EvioEasyNetPlugin struct {
	Conn net.Conn

	Ctx context.Context

	Config *YamlConfig

	Server *EvioServer

	Handler _interface.IEasyNet
}

func NewEvioEasyNetPlugin(ctx context.Context, iconfig _interface.IConfig, handler _interface.IEasyNet) *EvioEasyNetPlugin {

	var config *YamlConfig
	var ok bool
	if config, ok = iconfig.(*YamlConfig); !ok {
		logger.Error("evio yaml error \n")
	}

	evioEasyNetPlugin := &EvioEasyNetPlugin{
		Ctx:     ctx,
		Config:  config,
		Handler: handler,
	}

	server := NewEvioServer(ctx, config, handler)
	evioEasyNetPlugin.Server = server

	return evioEasyNetPlugin
}

func (g EvioEasyNetPlugin) Run() error {
	return g.Server.Run()
}
