package easynet

import (
	"io/ioutil"

	"github.com/baickl/logger"
	"github.com/EternalVow/easynet/interface"
	"github.com/EternalVow/easynet/plugin/evio"
	"github.com/EternalVow/easynet/plugin/gev"
	"github.com/EternalVow/easynet/plugin/gnet"
	"github.com/EternalVow/easynet/plugin/net"
	"github.com/EternalVow/easynet/plugin/netpoll"
	"gopkg.in/yaml.v2"
)

/*
	{
		"protocol":"tcp",
		"ip":"127.0.0.1",
		"port":80
	}
*/

type YamlAllConfig struct {
	EvioConfig    *evio.YamlConfig    `json:"evio_config" yaml:"evio_config"`
	GevConfig     *gev.YamlConfig     `json:"gev_config" yaml:"gev_config"`
	GnetConfig    *gnet.YamlConfig    `json:"gnet_config" yaml:"gnet_config"`
	NetConfig     *net.YamlConfig     `json:"net_config" yaml:"net_config"`
	NetpollConfig *netpoll.YamlConfig `json:"netpoll_config" yaml:"netpoll_config"`
}

type DeFaultNetConfig struct {
	Protocol string `json:"protocol"`
	Ip       string `json:"ip"`
	Port     int32  `json:"port"`
}

func NewDefaultNetConfig(Protocol string, Ip string, Port int32) _interface.IConfig {
	return &DeFaultNetConfig{
		Protocol: Protocol,
		Ip:       Ip,
		Port:     Port,
	}
}

// todo yaml
func NewNetConfigWithConfig(path string, netName string) _interface.IConfig {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		logger.Errorf("read yamlFile err :%v", err)
	}
	//将配置文件读取到结构体中
	yamlAllConfig := &YamlAllConfig{}
	err = yaml.Unmarshal(yamlFile, yamlAllConfig)
	if err != nil {
		logger.Errorf(" yamlFile Unmarshal err :%v", err)
	}
	logger.Infoln("read yamlFile yamlAllConfig :%v", yamlAllConfig)

	//var _config *config.Config
	var config _interface.IConfig
	switch netName {
	case "Gnet":

		config = yamlAllConfig.GnetConfig
	case "Gev":
		config = yamlAllConfig.GevConfig
	case "Net":
		config = yamlAllConfig.NetConfig
	case "NetPoll":
		config = yamlAllConfig.NetpollConfig
	case "Evio":
		config = yamlAllConfig.EvioConfig

	default:
		logger.Errorln("no expected net name")
	}
	logger.Infoln("read yamlFile :%v", config)

	return config
}

func (n *DeFaultNetConfig) GetProtocol() string {
	return n.Protocol
}

func (n *DeFaultNetConfig) GetIp() string {
	return n.Ip
}

func (n *DeFaultNetConfig) GetPort() int32 {
	return n.Port
}
