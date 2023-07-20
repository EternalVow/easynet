package gev

type YamlConfig struct {
	Protocol  string `json:"protocol" yaml:"protocol"`
	Ip        string `json:"ip" yaml:"ip"`
	Port      int32  `json:"port" yaml:"port"`
	Numloops  int32  `json:"numloops" yaml:"numloops"`
	Reuseport bool   `json:"reuseport" yaml:"reuseport"`
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

func (n *YamlConfig) GetNumloops() int32 {
	return n.Numloops
}

func (n *YamlConfig) GetReuseport() bool {
	return n.Reuseport
}
