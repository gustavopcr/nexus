package peer

var peers = make([]Peer, 0)

type Peer struct {
	PeerId    string `json:"peerId"`
	IPAddress string `json:"ipAddress"`
	Port      uint16 `json:"port"`
}

func NewPeer(p Peer) {
	peers = append(peers, p)
}

func GetAllPeers() []Peer {
	return peers
}
