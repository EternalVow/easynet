package net

type YamlConfig struct {
	Protocol string `json:"protocol" yaml:"protocol"`
	Ip       string `json:"ip" yaml:"ip"`
	Port     int32  `json:"port" yaml:"port"`
}

func (n *YamlConfig) GetProtocol() string {
	return n.Protocol
}

func (n *YamlConfig) GetIp() string {
	return n.Ip
}

func (n *YamlConfig) GetPort() int32 {
	return n.Port
}
