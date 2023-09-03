package socket

import (
	"Receiver/customCache"
	"context"
	"fmt"
	"sync/atomic"
)

var flag int32 = 0

func Check() {
	if atomic.LoadInt32(&flag) == 1 {
		return
	}
	atomic.StoreInt32(&flag, 1)
	defer atomic.StoreInt32(&flag, 0)

	message := customCache.Message{}
	messages, _ := message.ReadAll(context.Background(), customCache.Db, "messages")

	err := CreateConnection()
	if err != nil {
		return
	}
	defer CloseConnection()

	for _, element := range messages {
		err := SendData([]byte(element.Message))
		if err != nil {
			fmt.Println("Error sending data:", err)
			return
		}
		element.Received = true
		err = element.Update(context.Background(), customCache.Db, "messages", element.Sequence)
		if err != nil {
			return
		}
	}
}
