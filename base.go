package mysmartrpc

import "context"

type (
	RpcFn func(ctx context.Context, args any, reply any) error

	Server interface {
		Serve() error
		RegisterStartFn(fn func())
		RegisterShutdown(fn func())
		RegisterObj(obj any) error
		RegisterFn(path string, fn RpcFn) error
	}

	ClientPool interface {
	}
)

const (
	RpcTcpProtocol  string = "tcp"
	RpcHttpProtocol string = "http"
	RpcQuicProtocol string = "quic"
	RpcKcpProtocol  string = "kcp"
	RpcUnixProtocol string = "unix"

	RpcDefProtocol string = RpcHttpProtocol
)
