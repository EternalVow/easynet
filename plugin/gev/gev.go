package gev

import (
	"context"
	"fmt"
	"github.com/Allenxuxu/gev"
	"github.com/baickl/logger"
	"github.com/EternalVow/easynet/base"
	"github.com/EternalVow/easynet/interface"
)

type GevServer struct {
	Ctx  context.Context
	addr string

	handler        _interface.IEasyNet
	InputStreamMap map[string]_interface.IInputStream
	ConnectioonMap map[string]_interface.IConnection
}

func NewGevServer(ctx context.Context, config *YamlConfig, handler _interface.IEasyNet) *GevServer {
	return &GevServer{
		Ctx:            ctx,
		addr:           fmt.Sprintf("%s://%s:%d", config.GetProtocol(), config.GetIp(), config.GetPort()),
		handler:        handler,
		InputStreamMap: make(map[string]_interface.IInputStream),
		ConnectioonMap: make(map[string]_interface.IConnection),
	}
}

func (s *GevServer) OnConnect(c *gev.Connection) {
	logger.Infoln("Gev server OnConnect ")
	s.InputStreamMap[c.PeerAddr()] = &base.InputStream{}
	s.ConnectioonMap[c.PeerAddr()] = &Connection{
		Conn: c,
	}
	err := s.handler.OnConnect(s.ConnectioonMap[c.PeerAddr()])
	if err != nil {
		logger.Errorf("Gnet OnConnect err %v\n", err)
	}
	return
}

func (s *GevServer) OnMessage(c *gev.Connection, ctx interface{}, data []byte) (out interface{}) {
	s.InputStreamMap[c.PeerAddr()].Begin(data)
	data, err := s.handler.OnReceive(s.ConnectioonMap[c.PeerAddr()], s.InputStreamMap[c.PeerAddr()])
	if err != nil {
		logger.Errorf("Gnet OnMessage err %v\n", err)
	}
	return data
}

func (s *GevServer) OnClose(c *gev.Connection) {
	s.InputStreamMap[c.PeerAddr()] = nil
	err := s.handler.OnClose(s.ConnectioonMap[c.PeerAddr()], nil)
	if err != nil {
		logger.Errorf("Gnet OnClose err %v\n", err)
	}
	return
}
