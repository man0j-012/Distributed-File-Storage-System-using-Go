package p2p

import (
	"fmt"
	"net"  // standard go lib for network prorgamming (communications, addresses) ...
	"sync" // go lib for synchronization primitives for concurrent programming protecting shared resources from deadlocks
)

//Listening for incoming connections, storing connected peers, and ensuring thread-safe access to the peer list
type TCPTransport struct{
	listenAddress string
	listener net.Listener

	mu sync.RWMutex
	peers map[net.Addr]Peer
}

//constructor function for tcptransport
func NewTCPTransport(listenAddr string) *TCPTransport{
	return &TCPTransport{
		listenAddress: listenAddr,
	}
}


//function is responsible for starting the TCP listener and initiating the process of accepting incoming connections
func (t *TCPTransport) ListenAndAccept() error{
var err error
t.listener, err = net.Listen("tcp","t.listenAddress")
if err!=nil{
	return err;
}
go t.StartAcceptLoop()
return nil
}

//function runs in a loop, continuously accepting incoming connections
func (t* TCPTransport) StartAcceptLoop(){
	for{
		conn, err := t.listener.Accept()
		if err!=nil{
			fmt.Printf("TCP accept error: %s\n",err)
			continue
		}
		go t.handleConn(conn)
	}
}

func (t* TCPTransport) handleConn(conn net.Conn){
fmt.Printf("new incoming connection %+v\n",conn,
)
}