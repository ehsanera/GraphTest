package socket

import (
	"net"
)

var conn net.Conn

func CreateConnection() error {
	var err error
	conn, err = net.Dial("tcp", "127.0.0.1:8081")
	if err != nil {
		return err
	}
	return nil
}

func SendData(data []byte) error {
	_, err := conn.Write(data)
	return err
}

func CloseConnection() {
	err := conn.Close()
	if err != nil {
		return
	}
}
