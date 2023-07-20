package gnet

type YamlConfig struct {
	Protocol       string `json:"protocol" yaml:"protocol"`
	Ip             string `json:"ip" yaml:"ip"`
	Port           int32  `json:"port" yaml:"port"`
	Multicore      bool   `json:"multicore" yaml:"multicore"`
	Lockosthread   bool   `json:"lockosthread" yaml:"lockosthread"`
	ReadBufferCap  int32  `json:"readbuffercap" yaml:"readbuffercap"`
	WriteBufferCap int32  `json:"writebuffercap" yaml:"writebuffercap"`
	NumEventLoop   int32  `json:"numeventloop" yaml:"numeventloop"`
	ReusePort      bool   `json:"reuseport" yaml:"reuseport"`
	ReuseAddr      bool   `json:"reuseaddr" yaml:"reuseaddr"`
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

func (n *YamlConfig) GetMulticore() bool {
	return n.Multicore
}

func (n *YamlConfig) GetLockosthread() bool {
	return n.Lockosthread
}

func (n *YamlConfig) GetReadBufferCap() int32 {
	return n.ReadBufferCap
}

func (n *YamlConfig) GetWriteBufferCap() int32 {
	return n.WriteBufferCap
}

func (n *YamlConfig) GetNumEventLoop() int32 {
	return n.NumEventLoop
}

func (n *YamlConfig) GetReusePort() bool {
	return n.ReusePort
}

func (n *YamlConfig) GetReuseAddr() bool {
	return n.ReuseAddr
}
