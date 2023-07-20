package evio

import (
	"context"
	"fmt"
	"log"
	"strconv"

	"github.com/baickl/logger"
	"github.com/EternalVow/easynet/base"
	"github.com/EternalVow/easynet/interface"
	"github.com/tidwall/evio"
)

type EvioServer struct {
	Ctx context.Context

	config *YamlConfig
	addr   string

	handler _interface.IEasyNet

	InputStreamMap map[string]_interface.IInputStream
	ConnectioonMap map[string]_interface.IConnection
}

func NewEvioServer(ctx context.Context, config *YamlConfig, handler _interface.IEasyNet) *EvioServer {

	s := &EvioServer{
		Ctx:            ctx,
		handler:        handler,
		config:         config,
		InputStreamMap: make(map[string]_interface.IInputStream),
		ConnectioonMap: make(map[string]_interface.IConnection),
	}
	if s.config != nil {
		s.addr = s.getAddr()
	}
	return s
}

func (s EvioServer) Run() error {
	var events evio.Events
	events.NumLoops = int(s.config.GetLoops())

	events.Serving = func(srv evio.Server) (action evio.Action) {
		logger.Infoln("evio server OnStart")
		err := s.handler.OnStart(nil)
		if err != nil {
			logger.Errorf("evio server OnStart error %v", err)
		}
		return
	}

	events.Opened = func(c evio.Conn) (out []byte, opts evio.Options, action evio.Action) {
		s.InputStreamMap[strconv.Itoa(c.AddrIndex())] = &base.InputStream{}
		s.ConnectioonMap[strconv.Itoa(c.AddrIndex())] = &Connection{
			Conn: c,
		}
		logger.Infoln("evio Opened OnConnect")
		err := s.handler.OnConnect(s.ConnectioonMap[strconv.Itoa(c.AddrIndex())])
		if err != nil {
			log.Printf("evio server OnConnect error %v", err)
		}
		return
	}

	events.Data = func(c evio.Conn, in []byte) (out []byte, action evio.Action) {
		s.InputStreamMap[strconv.Itoa(c.AddrIndex())].Begin(in)
		out, err := s.handler.OnReceive(s.ConnectioonMap[strconv.Itoa(c.AddrIndex())], s.InputStreamMap[strconv.Itoa(c.AddrIndex())])
		if err != nil {
			logger.Errorf("evio server OnReceive err %v", err)
		}
		return
	}
	events.Closed = func(c evio.Conn, inErr error) (action evio.Action) {
		s.InputStreamMap[strconv.Itoa(c.AddrIndex())] = nil
		logger.Infoln("evio Opened OnClose")
		err := s.handler.OnClose(s.ConnectioonMap[strconv.Itoa(c.AddrIndex())], inErr)
		if err != nil {
			logger.Errorf("evio server OnClose error %v", err)
		}
		return
	}
	err := evio.Serve(events, s.addr)
	if err != nil {
		logger.Errorf("evio Serve error %v", err)
		return err
	}
	return nil
}

func (s EvioServer) getAddr() string {
	if s.config.GetStdlib() {
		ssuf := "-net"
		return fmt.Sprintf("%s%s://%s:%d?reuseport=%t", s.config.GetProtocol(), ssuf, s.config.GetIp(), s.config.GetPort(), s.config.GetReuseport())
	}
	return fmt.Sprintf("%s://%s:%d?reuseport=%t", s.config.GetProtocol(), s.config.GetIp(), s.config.GetPort(), s.config.GetReuseport())
}
