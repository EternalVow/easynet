package net

import (
	"context"
	"net"

	"github.com/baickl/logger"
	"github.com/EternalVow/easynet/interface"
)

type NetEasyNetPlugin struct {
	Conn net.Conn

	Ctx context.Context

	Config *YamlConfig

	Server *NetServer

	Handler _interface.IEasyNet
}

func NewNetEasyNetPlugin(ctx context.Context, iconfig _interface.IConfig, handler _interface.IEasyNet) *NetEasyNetPlugin {

	var config *YamlConfig
	var ok bool
	if config, ok = iconfig.(*YamlConfig); !ok {
		logger.Errorln("net yaml error ")
	}

	easyNetPlugin := &NetEasyNetPlugin{
		Ctx:     ctx,
		Config:  config,
		Handler: handler,
	}

	Server := NewNetServer(ctx, config, handler)
	easyNetPlugin.Server = Server

	return easyNetPlugin
}

func (g NetEasyNetPlugin) Run() error {
	return g.Server.Run()
}
