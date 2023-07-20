# Easynet

## 介绍
鉴于目前市面上存在各种各样的go网络库，不同的网络库各有不同的优势和劣势，当我们使用网络库实现应用层协议，往往我们只能根据选择其中一个网络库进行接入，假如业务产生变化，切换网络的成本比较巨大。于此，对市面上常用的网络库进行了一层封装，透出统一的接口，构建了全新的网络库——easynet。用户使用easynet，可以通过yaml配置选择不同的网络库，让切换网络库成本几乎为0。

_github.com/EternalVow/easynet_

## 架构说明

目前通过插件设计的模式，接入了5个市面上比较受欢迎的网络库，如下

- Gnet
- Gev
- Net(原生)
- NetPoll
- Evio

当我们使用这些网络库的时候，不需要理解里面的底层，easynet通过IEasyNet接口定义了统一的接入方式；
用户使用只需要通过yaml配置文件定义网络参数，入口传入网络库名称，定义继承IEasyNet接口实现，即可实现业务逻辑

```go
type IEasyNet interface {
	OnStart(conn interface{}) error

	OnConnect(conn interface{}) error

	OnReceive(conn interface{}, bytes []byte) ([]byte, error)

	OnShutdown(conn interface{}) error

	OnClose(conn interface{}, err error) error

	// todo to add more
}

```

目录架构如下

```text
├── base
│   ├── confg.yaml  配置文件示例
│   └── config.go   config 结构定义
├── example     示例
│   └── echo
│       ├── client.go 
│       └── server.go
├── interface  接口定义
│   ├── config.go
│   ├── easynet.go
│   └── plugin.go
└── plugin  插件目录，接入网络库核心代码
    ├── evio
    │   ├── evio.go 对接网络库代码
    │   ├── evio_plugin.go 插件定义结构体
    │   └── yaml_config.go 配置文件定义结构体
    ├── gev
    │   ├── gev.go
    │   ├── gev_plugin.go
    │   └── yaml_config.go
    ├── gnet
    │   ├── gnet.go
    │   ├── gnet_plugin.go
    │   └── yaml_config.go
    ├── net
    │   ├── net.go
    │   ├── net_plugin.go
    │   └── yaml_config.go
    └── netpoll
        ├── netpoll.go
        ├── netpoll_plugin.go
        └── yaml_config.go
├── easynet.go easynet网络库入口结构体
├── easynet_test.go
├── go.mod
├── go.sum
├── info.puml
```

## 使用说明

### 1.引入包
    
    `go get github.com/EternalVow/easynet`

### 2.定义handle 结构体

**目前handle结构体正常定义5种方法 OnStart/OnConnect/OnReceive/OnShutdown/OnClose**

```go



type Handler struct {
}

func (h Handler) OnStart(conn interface{}) error {
	fmt.Println("test OnStart")
	return nil
}

func (h Handler) OnConnect(conn interface{}) error {
	netpollConn, ok:= conn.(netpoll.Connection)
	if !ok {
		fmt.Println("test conn err")
	}

	fmt.Println("test conn LocalAddr",netpollConn.LocalAddr())
	fmt.Println("test conn RemoteAddr",netpollConn.RemoteAddr())

	return nil

}

func (h Handler) OnReceive(conn interface{}, bytes []byte) ([]byte, error) {

	fmt.Println("test receive msg ",string(bytes))

	return bytes, nil

}

func (h Handler) OnShutdown(conn interface{}) error {
	return nil
}

func (h Handler) OnClose(conn interface{}, err error) error {
	return nil
}

```

### 3.启动server

启动server 可以使用 `NewEasyNetWithYamlConfig`
netname 参数可以支持其中5个

- Gnet
- Gev
- Net
- NetPoll
- Evio

```go
func main() {
	handler := &Handler{}
	easynet.NewEasyNetWithYamlConfig(context.Background(), "NetPoll", handler, "../../base/confg.yaml")
}
```

其中涉及一个yaml配置文件,可以直接定义不同网络库的参数,示例如下

```yaml
evio_config:
  protocol: tcp
  ip: 127.0.0.1
  port: 9011
  reuseport: false
  stdlib: false
#  可以选择使用 stdlib（stdlib 主要是为了支持 非 *unix 平台）
  loops: 1
gev_config:
  protocol: tcp
  ip: 127.0.0.1
  port: 9011
  numloops: 1
  reuseport: false
gnet_config:
  protocol: tcp
  ip: 127.0.0.1
  port: 9011
  multicore: true
  lockosthread: false
  readbuffercap: 1000
  writebuffercap: 1000
  numeventloop: 100
  reuseport: false
  reuseaddr: false
net_config:
  protocol: tcp
  ip: 127.0.0.1
  port: 9011
netpoll_config:
  protocol: tcp
  ip: 127.0.0.1
  port: 9011

log:
  console: true
  level: debug
  dir: ./log

```






