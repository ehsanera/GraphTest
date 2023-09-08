package main

import (
	"Broker/customCache"
	"Broker/socket"
)

func main() {
	client, err := customCache.Connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	customCache.Db = client.Database("broker")

	socket.ServerConnect()
}
