package socket

import "net"
import "fmt"
import "os"

func Client(bytes []byte) {
	// Server address and port
	serverAddr := "127.0.0.1:8081"

	// Connect to the server
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	// Send the message bytes to the server
	_, err = conn.Write(bytes)
	if err != nil {
		fmt.Println("Error sending data:", err)
		os.Exit(1)
	}
}
