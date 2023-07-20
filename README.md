# Easynet
## Introduction
Given the existence of various Go network libraries on the market, each with its own advantages and disadvantages, when we use network libraries to implement application layer protocols, we often have to choose one of them for access. If there is a change in business, the cost of switching networks is relatively high. Here, a layer of encapsulation has been applied to commonly used network libraries on the market, revealing a unified interface, and a brand new network library - **easynet** has been constructed. Users using **easynet** can choose different network libraries through yaml configuration, making the cost of switching network libraries almost zero.

github.com/EternalVow/easynet

## Architecture Description
Currently, through the plugin design mode, we have accessed 5 popular network libraries on the market, as follows
-Gnet
-Gev
-Net (native)
-NetPoll
-Evio
When we use these network libraries, we don't need to understand the underlying layers. Easynet defines a unified access method through the IEasyNet interface;
Users only need to define network parameters through the yaml configuration file, pass in the network library name at the entrance, and inherit the IEasyNet interface implementation to implement business logic
```Go
type IEasyNet interface{
OnStart (conn interface {}) error
OnConnect (conn interface {}) error
OnReceive (conn interface {}, bytes [] byte) ([] byte, error)
On Shutdown (conn interface {}) error
OnClose (conn interface {}, err error) error
//Todo to add more
}
```
The directory structure is as follows


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

## Instructions for use
### 1. Import Package
`Go get github. com/EternalVow/easynet`
### 2. Define the handle structure
**Currently, there are 5 methods defined for the handle structure: OnStart/OnConnect/OnReceive/OnShutdown/Close**



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
### 3. Start the server
Starting the server can use 'NewEasyNetWithYamlConfig'`

The netname parameter can support 5 of them

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
It involves a yaml configuration file that can directly define parameters for different network libraries, as shown in the following example

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
