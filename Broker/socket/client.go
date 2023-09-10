package socket

import (
	"encoding/binary"
	"errors"
	"net"
)

func SendData(data []byte) error {
	connection, err := net.Dial("tcp", "127.0.0.1:8082")
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
	if err != nil {
		return err
	}

	if size == int64(len(data)) {
		return nil
	} else {
		return errors.New("oops")
	}
}
