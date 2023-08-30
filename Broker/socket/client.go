package socket

import (
	"net"
	"time"
)

var conn net.Conn

func createConnection(serverAddr string) (net.Conn, error) {
	// Connect to the server
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func isConnActive(conn net.Conn) bool {
	// Set a short deadline for testing connection status
	deadline := time.Now().Add(time.Second)

	// Try to read a byte with the given deadline
	err := conn.SetReadDeadline(deadline)
	if err != nil {
		return false
	}

	buffer := make([]byte, 1)
	_, err = conn.Read(buffer)
	if err != nil {
		return false
	}

	// Reset the deadline
	err = conn.SetReadDeadline(time.Time{})
	if err != nil {
		return false
	}
	return true
}

func SendData(data []byte) error {
	if conn == nil || !isConnActive(conn) {
		// Connection not established or not active, attempt to create it
		var err error
		conn, err = createConnection("127.0.0.1:8081")
		if err != nil {
			return err
		}
	}

	// Send the message bytes to the server
	_, err := conn.Write(data)
	return err
}
