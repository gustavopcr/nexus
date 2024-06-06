package routes

import (
	"fmt"
	"net/http"
	"net"
	"strconv"
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
		ip, portStr, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil{
			fmt.Println("err: ", err)
			http.Error(w, "Invalid Ip", http.StatusForbidden)
			return
		}
		port, err := strconv.ParseUint(portStr, 10, 16)
		if err != nil{
			fmt.Println("err: ", err)
			http.Error(w, "Invalid Port", http.StatusForbidden)
			return
		}
		peer.NewPeer( peer.Peer{ PeerId: "alo123", IPAddress: ip, Port: uint16(port)})
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("/peers", handleOnConnectPeer)

	return router
}
