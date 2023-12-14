package mysmartrpc

import (
	"fmt"
	"github.com/smallnest/rpcx/server"
)

type rpcServer struct {
	*server.Server
	protocol string
	port     int

	cbOnRestart  func()
	cbOnShutdown func()
}

func NewServer(protocol string, port int, option ...server.OptionFn) Server {
	s := &rpcServer{port: port}
	switch protocol {
	case RpcTcpProtocol, RpcHttpProtocol, RpcKcpProtocol, RpcQuicProtocol, RpcUnixProtocol:
		s.protocol = protocol
		break
	default:
		s.protocol = RpcDefProtocol
		break
	}

	s.Server = server.NewServer(option...)
	return s
}

func (self *rpcServer) Serve() error {
	return self.Server.Serve(self.protocol, fmt.Sprintf(":%d", self.port))
}

func (self *rpcServer) RegisterStartFn(fn func()) {
	self.cbOnRestart = fn
	self.Server.RegisterOnRestart(self.onReStart)
}

func (self *rpcServer) RegisterShutdown(fn func()) {
	self.cbOnShutdown = fn
	self.Server.RegisterOnShutdown(self.onShutdown)
}

func (self *rpcServer) RegisterObj(obj any) error {
	return self.Server.Register(obj, "")
}

func (self *rpcServer) RegisterFn(path string, fn RpcFn) error {
	return self.Server.RegisterFunction(path, fn, "")
}

func (self *rpcServer) onReStart(s *server.Server) {
	if self.cbOnRestart != nil {
		self.cbOnRestart()
	}
}

func (self *rpcServer) onShutdown(s *server.Server) {
	if self.cbOnShutdown != nil {
		self.cbOnShutdown()
	}
}
