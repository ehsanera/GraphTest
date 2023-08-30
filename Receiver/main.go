package main

import (
	"Receiver/customCache"
	"Receiver/server"
)

func main() {
	client, err := customCache.Connect("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	customCache.Db = client.Database("receiver")

	server.StartServer()
}
