package evio

type YamlConfig struct {
	Protocol  string `json:"protocol" yaml:"protocol"`
	Ip        string `json:"ip" yaml:"ip"`
	Port      int32  `json:"port" yaml:"port"`
	Reuseport bool   `json:"reuseport" yaml:"reuseport"`
	Stdlib    bool   `json:"stdlib" yaml:"stdlib"`
	Loops     int32  `json:"loops" yaml:"loops"`
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

func (n *YamlConfig) GetReuseport() bool {
	return n.Reuseport
}

func (n *YamlConfig) GetLoops() int32 {
	return n.Loops
}

func (n *YamlConfig) GetStdlib() bool {
	return n.Stdlib
}
