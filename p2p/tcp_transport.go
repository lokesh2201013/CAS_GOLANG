package p2p

import (
	"fmt"
	"net"
	"sync"
)

type TCPPeer struct{
	conn net.Conn
	outbound bool
}
type TCPTransport struct {
	listenAddress      string
	peers              map[net.Addr ]Peer
	handshakeFunc      HandshakeFunc
	decoder            Decoder
	listener           net.Listener
	mu                 sync.RWMutex
 
}

func NewTCPPeer(conn net.Conn,outbound bool) *TCPPeer{
return &TCPPeer{
	conn: conn,
	outbound:outbound,
}
}

func NewTCPTransport(listenAddr string)* TCPTransport{
	return &TCPTransport{
		handshakeFunc: NOPHandshakeFunc,
		listenAddress: listenAddr,
        
	}
}


func (t *TCPTransport)ListenAndAccept() error{
	var err error
	t.listener,err=net.Listen("tcp",t.listenAddress)
	if err != nil {
	return err
	}
	go t.startAcceptLoop()
	return nil
}

func (t *TCPTransport)startAcceptLoop(){
	for{
		conn, err := t.listener.Accept()
		if err != nil {
			fmt.Println("TCP accept error:%s\n",err)
		}
		go t.handleConn(conn)
	}
}

func (t *TCPTransport) handleConn(conn net.Conn){
	peer:= NewTCPPeer(conn,true)
	if err:=t.handshakeFunc(conn);err!=nil{
		  
	}
	fmt.Println("new incoming connection %+v\n",peer )
}