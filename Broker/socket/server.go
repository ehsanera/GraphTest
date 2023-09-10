package socket

import (
	"Broker/customCache"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
)

func ServerConnect() {
	listenAddr := "127.0.0.1:8081"

	server, err := net.Listen("tcp", listenAddr)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	defer server.Close()
	fmt.Println("Listening on 127.0.0.1:8081...")
	for {
		connection, err := server.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		go processClient(connection)
	}
}

func processClient(connection net.Conn) {
	buffer := make([]byte, 8196)
	mLen, err := connection.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}

	message := customCache.Message{}
	err = json.Unmarshal(buffer[:mLen], &message)
	if err != nil {
		return
	}

	err = message.Create(context.Background(), customCache.Db, "messages", message)
	if err != nil {
		return
	}
	log.Println("Received: ", mLen)
	log.Println("Message: ", message)
	connection.Close()
	go Send()
}
