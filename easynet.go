package easynet

import (
	"context"
	"fmt"
	"github.com/baickl/logger"
	"github.com/EternalVow/easynet/interface"
	"github.com/EternalVow/easynet/plugin/evio"
	"github.com/EternalVow/easynet/plugin/gev"
	"github.com/EternalVow/easynet/plugin/gnet"
	np "github.com/EternalVow/easynet/plugin/net"
	"github.com/EternalVow/easynet/plugin/netpoll"
	"net"
	"unsafe"
)

func init() {
	logger.Initialize("./log", "LoginServer")
}

type EasyNet struct {
	handler _interface.IEasyNet

	Conn net.Conn

	Ctx context.Context

	EasyNetPlugin _interface.IPlugin

	Config _interface.IConfig
}

func NewEasyNet(ctx context.Context, netName string, config _interface.IConfig, handler _interface.IEasyNet) *EasyNet {
	easynet := &EasyNet{
		Ctx:     ctx,
		handler: handler,
		Config:  config,
	}

	var configImpl _interface.IConfig = config

	var defaultConfigImpl *DeFaultNetConfig
	if dc, ok := config.(*DeFaultNetConfig); ok {
		defaultConfigImpl = dc
	}

	// todo new EasyNetPlugin
	switch netName {
	case "Gnet":
		if defaultConfigImpl != nil {
			configImpl = (*gnet.YamlConfig)(unsafe.Pointer(defaultConfigImpl))
		}
		easynet.EasyNetPlugin = gnet.NewGnetEasyNetPlugin(ctx, configImpl, handler)
		err := easynet.EasyNetPlugin.Run()
		if err != nil {
			fmt.Println(err)
		}
	case "Gev":
		if defaultConfigImpl != nil {
			configImpl = (*gev.YamlConfig)(unsafe.Pointer(defaultConfigImpl))
		}
		easynet.EasyNetPlugin = gev.NewGevEasyNetPlugin(ctx, configImpl, handler)
		err := easynet.EasyNetPlugin.Run()
		if err != nil {
			fmt.Println(err)
		}
	case "Net":
		if defaultConfigImpl != nil {
			configImpl = (*np.YamlConfig)(unsafe.Pointer(defaultConfigImpl))
		}
		easynet.EasyNetPlugin = np.NewNetEasyNetPlugin(ctx, configImpl, handler)
		err := easynet.EasyNetPlugin.Run()
		if err != nil {
			fmt.Println(err)
		}
	case "NetPoll":
		if defaultConfigImpl != nil {
			configImpl = (*netpoll.YamlConfig)(unsafe.Pointer(defaultConfigImpl))
		}
		easynet.EasyNetPlugin = netpoll.NewNetPollEasyNetPlugin(ctx, configImpl, handler)
		err := easynet.EasyNetPlugin.Run()
		if err != nil {
			fmt.Println(err)
		}
	case "Evio":
		if defaultConfigImpl != nil {
			configImpl = (*evio.YamlConfig)(unsafe.Pointer(defaultConfigImpl))
		}
		easynet.EasyNetPlugin = evio.NewEvioEasyNetPlugin(ctx, configImpl, handler)
		err := easynet.EasyNetPlugin.Run()
		if err != nil {
			fmt.Println(err)
		}
	default:
		logger.Errorln("no expected net plugin")
	}

	return easynet
}

func NewEasyNetWithYamlConfig(ctx context.Context, netName string, handler _interface.IEasyNet, path string) *EasyNet {
	easynet := &EasyNet{
		Ctx:     ctx,
		handler: handler,
	}

	// todo new EasyNetPlugin
	switch netName {
	case "Gnet":
		config := NewNetConfigWithConfig(path, netName)
		easynet.Config = config
		easynet.EasyNetPlugin = gnet.NewGnetEasyNetPlugin(ctx, config, handler)
		err := easynet.EasyNetPlugin.Run()
		if err != nil {
			logger.Errorf("Gnet run err :%v", err)
		}
	case "Gev":
		config := NewNetConfigWithConfig(path, netName)
		easynet.Config = config
		easynet.EasyNetPlugin = gev.NewGevEasyNetPlugin(ctx, config, handler)
		err := easynet.EasyNetPlugin.Run()
		if err != nil {
			logger.Errorf("Gev run err :%v", err)
		}
	case "Net":
		config := NewNetConfigWithConfig(path, netName)
		easynet.Config = config
		easynet.EasyNetPlugin = np.NewNetEasyNetPlugin(ctx, config, handler)
		err := easynet.EasyNetPlugin.Run()
		if err != nil {
			logger.Errorf("Net run err :%v", err)
		}
	case "NetPoll":
		config := NewNetConfigWithConfig(path, netName)
		easynet.Config = config
		easynet.EasyNetPlugin = netpoll.NewNetPollEasyNetPlugin(ctx, config, handler)
		err := easynet.EasyNetPlugin.Run()
		if err != nil {
			logger.Errorf("NetPoll run err :%v", err)
		}
	case "Evio":
		config := NewNetConfigWithConfig(path, netName)
		easynet.Config = config
		easynet.EasyNetPlugin = evio.NewEvioEasyNetPlugin(ctx, config, handler)
		err := easynet.EasyNetPlugin.Run()
		if err != nil {
			logger.Errorf("Evio run err :%v", err)
		}
	default:
		logger.Errorln("no expected net plugin")
	}

	return easynet
}
