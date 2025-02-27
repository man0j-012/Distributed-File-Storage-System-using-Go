package p2p

// Peer represents the remote node in the network system or topology
type Peer interface{

}

//Transport handles the communication between nodes of the network. (TCP/IP, UDP, WebSockets)
type Transport interface{
	ListenAndAccept() error
}