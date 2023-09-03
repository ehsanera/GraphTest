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

	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server listening on", listenAddr)

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 8192)
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

	err = message.Create(context.Background(), customCache.Db, "messages", message)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Received %d\n", n)
}
