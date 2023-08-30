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

	message := customCache.Message{}

	var messages []customCache.Message

	err := message.ReadAll(context.Background(), customCache.Db, "messages", messages)
	if err != nil {
		return
	}

	for _, element := range messages {
		err = SendData([]byte(element.Message))
		if err != nil {
			fmt.Println("Error sending data:", err)
			return
		}
		err := element.Update(context.Background(), customCache.Db, "messages", message, message)
		if err != nil {
			return
		}
	}
}
