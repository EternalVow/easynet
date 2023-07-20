package net

import (
	"context"
	"fmt"
	"github.com/EternalVow/easynet/base"
	"net"
	"time"

	"github.com/baickl/logger"
	"github.com/EternalVow/easynet/interface"
)

type NetServer struct {
	Ctx context.Context

	Network string
	Address string

	handler        _interface.IEasyNet
	InputStreamMap map[string]_interface.IInputStream
	ConnectionMap  map[string]_interface.IConnection
}

func NewNetServer(ctx context.Context, config *YamlConfig, handler _interface.IEasyNet) *NetServer {
	return &NetServer{
		Ctx:            ctx,
		Network:        config.GetProtocol(),
		Address:        fmt.Sprintf("%s:%d", config.GetIp(), config.GetPort()),
		handler:        handler,
		InputStreamMap: make(map[string]_interface.IInputStream),
		ConnectionMap:  make(map[string]_interface.IConnection),
	}

}

func (s *NetServer) Run() error {
	ln, err := net.Listen(s.Network, s.Address)
	if err != nil {
		return err
	}
	if err := s.handler.OnStart(nil); err != nil {
		logger.Errorf("net OnStart err %v", err)
		return err
	}
	logger.Infoln("net OnStart")

	for {
		conn, err := ln.Accept()
		if err != nil {
			// handle error
			continue
		}
		s.InputStreamMap[conn.RemoteAddr().String()] = &base.InputStream{}
		s.ConnectionMap[conn.RemoteAddr().String()] = &Connection{
			Conn: conn,
		}
		if err = s.handler.OnConnect(s.ConnectionMap[conn.RemoteAddr().String()]); err != nil {
			// handle error
			continue
		}
		logger.Infoln("net OnConnect")

		go s.handleConnection(conn)
	}
	return nil
}

func (s *NetServer) handleConnection(conn net.Conn) {
	//函数调用完毕，自动关闭conn
	defer func(conn net.Conn) {
		conn.Close()
		s.InputStreamMap[conn.RemoteAddr().String()] = nil
		s.handler.OnClose(s.ConnectionMap[conn.RemoteAddr().String()], nil)
	}(conn)

	//4、获取客户端的网络地址信息
	rbuf, wbuf := make([]byte, 0xFFF), []byte{}

	//5、获取用户数据
	for {
		rlen, err := conn.Read(rbuf)
		if err != nil {
			logger.Errorf("net read message err %v", err)
			return
		}
		if rlen <= 0 {
			//logger.Infoln("rlen err %v", rlen)
			time.Sleep(time.Second * 1)
			continue
		}
		s.InputStreamMap[conn.RemoteAddr().String()].Begin(rbuf[:rlen])
		if wbuf, err = s.handler.OnReceive(s.ConnectionMap[conn.RemoteAddr().String()], s.InputStreamMap[conn.RemoteAddr().String()]); err != nil {
			logger.Errorf("net OnReceive err %v", err)
			return
		}
		//6、给用户发送回去
		if len(wbuf) > 0 {
			conn.Write(wbuf)
		}
	}
}
