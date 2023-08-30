package socket

import (
	"Broker/customCache"
	"context"
	"fmt"
	"log"
	"net"
)

func ServerConnect() {
	listenAddr := "127.0.0.1:8081"

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

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	receivedData := buffer[:n]

	message := customCache.Message{
		Message:  receivedData,
		Received: false,
	}

	err = message.Create(context.Background(), customCache.Db, "messages", &message)
	if err != nil {
		panic(err)
	}

	log.Printf("Received %d bytes: %s\n", n, string(receivedData))

	err = SendData(message.Message)
	if err != nil {
		return
	}
}
