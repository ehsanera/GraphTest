package socket

import (
	"fmt"
	"log"
	"net"
)

func ServerConnect() {
	listenAddr := "127.0.0.1:8082"

	// Listen for incoming connections
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server listening on", listenAddr)

	for {
		// Accept incoming connection
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// Handle connection in a separate goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 8192)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Printf("Error reading from %s: %v\n", conn.RemoteAddr(), err)
		return
	}

	log.Printf("Received %d\n", n)
}
