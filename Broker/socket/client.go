package socket

import (
	"encoding/binary"
	"errors"
	"net"
)

const (
	ServerHost = "localhost"
	ServerPort = "8082"
	ServerType = "tcp"
)

func SendData(data []byte) error {
	connection, err := net.Dial(ServerType, ServerHost+":"+ServerPort)
	if err != nil {
		return err
	}
	defer connection.Close()

	_, err = connection.Write(data)
	if err != nil {
		return err
	}

	var size int64
	err = binary.Read(connection, binary.LittleEndian, &size)
	if err == nil {
		return nil
	}

	return errors.New("oops")
}
