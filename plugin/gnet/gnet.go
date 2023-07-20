package gnet

import (
	"context"
	"fmt"
	"github.com/EternalVow/easynet/base"
	"github.com/baickl/logger"
	"strconv"
	"time"

	"github.com/EternalVow/easynet/interface"
	"github.com/panjf2000/gnet/v2"
	//"time"
)

type GnetServer struct {
	Ctx context.Context
	gnet.BuiltinEventEngine

	eng    gnet.Engine
	addr   string
	config *YamlConfig

	handler        _interface.IEasyNet
	InputStreamMap map[string]_interface.IInputStream
	ConnectionMap  map[string]_interface.IConnection
}

func NewGnetServer(ctx context.Context, config *YamlConfig, handler _interface.IEasyNet) *GnetServer {
	server := &GnetServer{
		Ctx:            ctx,
		handler:        handler,
		config:         config,
		InputStreamMap: make(map[string]_interface.IInputStream),
		ConnectionMap:  make(map[string]_interface.IConnection),
	}

	server.addr = server.getAddr()

	return server

}

func (s *GnetServer) OnBoot(eng gnet.Engine) gnet.Action {
	s.eng = eng
	logger.Infof("Gnet OnStart with multi-core=%t is listening on %s\n", s.config.Multicore, s.addr)
	s.handler.OnStart(nil)
	return gnet.None
}

func (s *GnetServer) OnShutdown(eng gnet.Engine) {
	logger.Infoln("Gnet OnShutdown")
	s.handler.OnShutdown(nil)
}

func (s *GnetServer) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	s.InputStreamMap[c.RemoteAddr().String()+strconv.Itoa(c.Fd())] = &base.InputStream{}
	s.ConnectionMap[c.RemoteAddr().String()+strconv.Itoa(c.Fd())] = &Connection{
		Conn: c,
	}

	logger.Infoln("Gnet OnConnect")
	s.handler.OnConnect(s.ConnectionMap[c.RemoteAddr().String()+strconv.Itoa(c.Fd())])

	return nil, gnet.None
}

func (s *GnetServer) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	s.InputStreamMap[c.RemoteAddr().String()+strconv.Itoa(c.Fd())] = nil
	logger.Infoln("Gnet OnClose")
	s.handler.OnClose(s.ConnectionMap[c.RemoteAddr().String()+strconv.Itoa(c.Fd())], err)

	return gnet.None
}

func (s *GnetServer) OnTick() (delay time.Duration, action gnet.Action) {
	logger.Infoln("Gnet OnTick")
	return 0, gnet.None
}

func (s *GnetServer) OnTraffic(c gnet.Conn) gnet.Action {
	logger.Infoln("Gnet OnReceive")
	// -a all buffer
	buf, _ := c.Next(-1)
	s.InputStreamMap[c.RemoteAddr().String()+strconv.Itoa(c.Fd())].Begin(buf)
	out, err := s.handler.OnReceive(s.ConnectionMap[c.RemoteAddr().String()+strconv.Itoa(c.Fd())], s.InputStreamMap[c.RemoteAddr().String()+strconv.Itoa(c.Fd())])
	if err != nil {
		return gnet.Close
	}
	if len(out) != 0 {
		_, err := c.Write(out)
		if err != nil {
			return gnet.Close
		}
	}
	return gnet.None
}

func (s GnetServer) getAddr() string {
	return fmt.Sprintf("%s://%s:%d", s.config.GetProtocol(), s.config.GetIp(), s.config.GetPort())
}
