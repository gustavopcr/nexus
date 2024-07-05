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

		p := peer.Peer {
			PeerId: "123",
			IPAddress: "1.1.1.1",
			Port: 6550,
		}
		/*
			PeerId    string `json:"peerId"`
			IPAddress string `json:"ipAddress"`
			Port      uint16 `json:"port"`
		
		*/
		peer.NewPeer(p)
		fmt.Printf("Received message from %s: %s\n", remoteAddr, string(buffer[:n]))

		// Send a response back to the client
		response := []byte("Message received!")
		_, err = conn.WriteToUDP(response, remoteAddr)
		if err != nil {
			fmt.Println("Error sending response:", err)
			return
		}
		fmt.Println("peers count: ", len(peer.GetAllPeers()))
		fmt.Printf("Sent response to %s: %s\n", remoteAddr, string(response))
	}
}
