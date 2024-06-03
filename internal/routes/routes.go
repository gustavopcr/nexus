package routes

import (
	"fmt"
	"net/http"

	"github.com/gustavopcr/nexus/internal/peer"
)

func handleGET(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "GET!")
}
func handleOnConnectPeer(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		fmt.Println("peers", peer.GetAllPeers())
		fmt.Fprintf(w, "alo peer")
	} else if r.Method == "POST" {
		fmt.Fprintf(w, "alo peer")
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", handleGET)
	router.HandleFunc("/peers", handleOnConnectPeer)

	return router
}
