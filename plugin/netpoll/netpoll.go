package netpoll

import (
	"context"
	"fmt"
	"github.com/EternalVow/easynet/base"

	"github.com/baickl/logger"
	"github.com/cloudwego/netpoll"
	"github.com/EternalVow/easynet/interface"
)

type NetPollServer struct {
	Ctx context.Context

	config *YamlConfig

	handler        _interface.IEasyNet
	InputStreamMap map[string]_interface.IInputStream
	ConnectionMap  map[string]_interface.IConnection
}

func NewNetPollServer(ctx context.Context, config *YamlConfig, handler _interface.IEasyNet) *NetPollServer {
	server := &NetPollServer{
		Ctx:            ctx,
		config:         config,
		handler:        handler,
		InputStreamMap: make(map[string]_interface.IInputStream),
		ConnectionMap:  make(map[string]_interface.IConnection),
	}
	return server

}

func (s *NetPollServer) Run() error {

	var eventLoop netpoll.EventLoop

	listener, err := netpoll.CreateListener(s.config.GetProtocol(), s.getAddr())
	if err != nil {
		logger.Errorf("create netpoll listener failed err:%v", err)
		return err
	}
	err = s.handler.OnStart(nil)
	if err != nil {
		logger.Errorf("create netpoll OnStart failed err:%v", err)
		return err
	}
	logger.Infof("create netpoll OnStart,Protocol:%v ,addr:%v", s.config.GetProtocol(), s.getAddr())

	//type OnRequest func(ctx context.Context, connection Connection) error
	handle := func(ctx context.Context, connection netpoll.Connection) error {

		var reader, writer = connection.Reader(), connection.Writer()

		// reading
		buf, _ := reader.Next(reader.Len())
		s.InputStreamMap[connection.RemoteAddr().String()].Begin(buf)
		reader.Release()
		//... parse the read data ...
		//var writeData []byte
		writeData, err := s.handler.OnReceive(s.ConnectionMap[connection.RemoteAddr().String()], s.InputStreamMap[connection.RemoteAddr().String()])
		if err != nil {
			logger.Errorf("netpoll server OnReceive ,err=$v \n", err)
			return err
		}

		// writing
		//... make the write data ...
		alloc, err := writer.Malloc(len(writeData))
		copy(alloc, writeData) // write data
		err = writer.Flush()
		if err != nil {
			logger.Errorf("netpoll server writing %s,err:%v \n", string(writeData), err)
		}
		logger.Infof("netpoll server WriteBinary %s,n:%v \n", string(writeData), len(writeData))
		return err
	}
	// todo is right
	//prepare := func(connection netpoll.Connection) context.Context {
	//	logger.Infoln("netpoll server OnStart")
	//	s.handler.OnStart(connection)
	//	return s.Ctx
	//}

	//type OnConnect func(ctx context.Context, connection Connection) context.Context
	connect := func(ctx context.Context, connection netpoll.Connection) context.Context {
		s.InputStreamMap[connection.RemoteAddr().String()] = &base.InputStream{}
		s.ConnectionMap[connection.RemoteAddr().String()] = &Connection{
			Conn: connection,
		}
		connection.AddCloseCallback(func(connection netpoll.Connection) error {
			s.InputStreamMap[connection.RemoteAddr().String()] = nil
			return nil
		})
		logger.Infoln("netpoll server OnConnect")
		err := s.handler.OnConnect(s.ConnectionMap[connection.RemoteAddr().String()])
		if err != nil {
			logger.Errorf("netpoll server OnConnect ,err=$v \n", err)
		}
		return ctx
	}

	eventLoop, _ = netpoll.NewEventLoop(
		handle,
		//netpoll.WithOnPrepare(prepare),
		netpoll.WithOnConnect(connect),
	)

	// start listen loop ...
	eventLoop.Serve(listener)

	return nil
}

func (s NetPollServer) getAddr() string {
	return fmt.Sprintf("%s:%d", s.config.GetIp(), s.config.GetPort())
}
