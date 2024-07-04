package routes

import (
	"fmt"
	"net/http"
	"github.com/gustavopcr/nexus/internal/peer"
	"encoding/json"
)

func handleGET(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GET!")
}
func handleOnConnectPeer(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		json.NewEncoder(w).Encode(peer.GetAllPeers())
		return
	} else if r.Method == "POST" {		
		peer.NewPeer( peer.Peer{ PeerId: "alo123", IPAddress: "10.0.0.1", Port: 6705})
		json.NewEncoder(w).Encode(peer.GetAllPeers())
		return
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/peer", handleOnConnectPeer)

	return router
}
