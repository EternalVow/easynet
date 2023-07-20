package gev

import (
	"context"
	"net"
	"strconv"

	"github.com/Allenxuxu/gev"
	"github.com/baickl/logger"
	"github.com/EternalVow/easynet/interface"
)

type GevEasyNetPlugin struct {
	Conn net.Conn

	Ctx context.Context

	Config *YamlConfig

	Server *GevServer

	Handler _interface.IEasyNet
}

func NewGevEasyNetPlugin(ctx context.Context, iconfig _interface.IConfig, handler _interface.IEasyNet) *GevEasyNetPlugin {

	var config *YamlConfig
	var ok bool
	if config, ok = iconfig.(*YamlConfig); !ok {
		logger.Errorln("gev yaml error ")
	}

	easyNetPlugin := &GevEasyNetPlugin{
		Ctx:     ctx,
		Config:  config,
		Handler: handler,
	}

	Server := NewGevServer(ctx, config, handler)
	easyNetPlugin.Server = Server

	return easyNetPlugin
}

func (g GevEasyNetPlugin) Run() error {

	var optionArr = []gev.Option{
		gev.Network(g.Config.Protocol),
		gev.Address(":" + (strconv.Itoa(int(g.Config.Port)))),
	}
	if g.Config.GetNumloops() != 0 {
		optionArr = append(optionArr, gev.NumLoops(int(g.Config.Numloops)))
	}
	if g.Config.GetReuseport() {
		optionArr = append(optionArr, gev.ReusePort(g.Config.Reuseport))
	}
	s, err := gev.NewServer(g.Server,
		optionArr...,
	)
	if err != nil {
		return err
	}

	s.Start()
	return nil
}
