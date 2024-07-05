package peer

import("net")

var peers = make([]Peer, 0)

type Peer struct {
	Address *net.UDPAddr
}

func NewPeer(p Peer) {
	peers = append(peers, p)
}

func GetAllPeers() []Peer {
	return peers
}

func ClearPeers(){
	peers = peers[:0]
}
