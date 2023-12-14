package mysmartrpc

import (
	"fmt"
	"github.com/smallnest/rpcx/client"
)

type rpcClient struct {
	*client.Peer2PeerDiscovery
	protocol   string
	serverAddr string
}

func NewClient(protocol, addr string) (*rpcClient, error) {
	c := &rpcClient{serverAddr: addr}
	switch protocol {
	case RpcTcpProtocol, RpcHttpProtocol, RpcKcpProtocol, RpcQuicProtocol, RpcUnixProtocol:
		c.protocol = protocol
		break
	default:
		c.protocol = RpcDefProtocol
		break
	}

	var err error
	c.Peer2PeerDiscovery, err = client.NewPeer2PeerDiscovery(fmt.Sprintf("%s@%s", c.protocol, c.serverAddr), "")
}
