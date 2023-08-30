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

	initialData := "Hello, client! Welcome to the server."
	_, err := conn.Write([]byte(initialData))
	if err != nil {
		log.Printf("Error sending initial data to %s: %v\n", conn.RemoteAddr(), err)
		return
	}

	log.Printf("Sent initial data to %s: %s\n", conn.RemoteAddr(), initialData)

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Printf("Error reading from %s: %v\n", conn.RemoteAddr(), err)
		return
	}

	receivedData := buffer[:n]
	log.Printf("Received %d bytes %s\n", n, string(receivedData))
}
