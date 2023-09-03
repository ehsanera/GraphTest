package socket

import (
	"net"
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

func SendData(data []byte) error {
	if conn == nil {
		// Connection not established or not active, attempt to create it
		var err error
		conn, err = createConnection("127.0.0.1:8082")
		if err != nil {
			return err
		}
	}

	// Send the message bytes to the server
	_, err := conn.Write(data)
	return err
}
