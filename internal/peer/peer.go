package peer

var peers = make([]Peer, 0)

type Peer struct {
	PeerId    string
	IPAddress string
	Port      uint16
}

func NewPeer(p Peer) {
	peers = append(peers, p)
}

func GetAllPeers() []Peer {
	return peers
}
