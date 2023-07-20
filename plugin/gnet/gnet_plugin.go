package gnet

import (
	"context"
	"gnet"
	"log"
	"net"

	"github.com/EternalVow/easynet/interface"
	"github.com/panjf2000/gnet/v2"
)

type GnetEasyNetPlugin struct {
	Conn net.Conn

	Ctx context.Context

	Config *YamlConfig

	Server *GnetServer

	Handler _interface.IEasyNet
}

func NewGnetEasyNetPlugin(ctx context.Context, iconfig _interface.IConfig, handler _interface.IEasyNet) *GnetEasyNetPlugin {

	var config *YamlConfig
	var ok bool
	if config, ok = iconfig.(*YamlConfig); !ok {
		log.Printf("gnet yaml error \n")
	}

	gnetEasyNetPlugin := &GnetEasyNetPlugin{
		Ctx:     ctx,
		Config:  config,
		Handler: handler,
	}

	gnetServer := NewGnetServer(ctx, config, handler)
	gnetEasyNetPlugin.Server = gnetServer

	return gnetEasyNetPlugin
}

func (g GnetEasyNetPlugin) Run() error {

	var optionArr = []gnet.Option{}

	if g.Config.GetMulticore() {
		optionArr = append(optionArr, gnet.WithMulticore(g.Config.GetMulticore()))
	}
	if g.Config.GetLockosthread() {
		optionArr = append(optionArr, gnet.WithLockOSThread(g.Config.GetLockosthread()))
	}
	if g.Config.GetReadBufferCap() > 0 {
		optionArr = append(optionArr, gnet.WithReadBufferCap(int(g.Config.GetReadBufferCap())))
	}
	if g.Config.GetWriteBufferCap() > 0 {
		optionArr = append(optionArr, gnet.WithWriteBufferCap(int(g.Config.GetWriteBufferCap())))
	}
	if g.Config.GetNumEventLoop() > 0 {
		optionArr = append(optionArr, gnet.WithNumEventLoop(int(g.Config.GetNumEventLoop())))
	}
	if g.Config.GetReusePort() {
		optionArr = append(optionArr, gnet.WithReusePort(g.Config.GetReusePort()))
	}
	if g.Config.GetReuseAddr() {
		optionArr = append(optionArr, gnet.WithReuseAddr(g.Config.GetReuseAddr()))
	}

	err := gnet.Run(
		g.Server,
		g.Server.addr,
		optionArr...)
	return err
}
