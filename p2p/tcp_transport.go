package p2p

import "net"

type TCPTransport struct {
	listenAddress string
	peers map[net.Addr ]
	listen        net.Listener
}