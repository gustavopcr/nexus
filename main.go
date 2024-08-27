package main

import (
	"fmt"
	"net"

	"github.com/gustavopcr/nexus/type/peer"
)

func main() {
	// Create a UDP address to listen on
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		fmt.Println("Error resolving address:", err)
		return
	}

	// Create a UDP connection to listen on
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening on address:", err)
		return
	}
	defer conn.Close()
	fmt.Println("Listening for UDP packets on port 8080...")

	buffer := make([]byte, 1024)

	for {
		// Read from the UDP connection
		n, remoteAddr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			fmt.Println("Error reading from UDP connection:", err)
			return
		}

		peer.NewPeer(peer.Peer{Address: remoteAddr})

		fmt.Printf("Received message from %s: %s\n", remoteAddr, string(buffer[:n]))
		fmt.Println("peers count: ", len(peer.GetAllPeers()))
		fmt.Println("peers count: ", peer.GetAllPeers())

		//fmt.Printf("Sent response to %s: %s\n", remoteAddr, string(response))

		if len(peer.GetAllPeers()) >= 2 {
			peers := peer.GetAllPeers()
			// Send each client the other's address information
			for i := 0; i < len(peers); i++ {
				for j := 0; j < len(peers); j++ {
					if i != j {
						message := fmt.Sprintf(peers[j].Address.String())
						_, err = conn.WriteToUDP([]byte(message), peers[i].Address)
						if err != nil {
							fmt.Println("Error sending peer info:", err)
						}
					}
				}
			}

			peer.ClearPeers()
		}

	}
}
