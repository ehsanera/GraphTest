package socket

import (
	"Broker/customCache"
	"context"
	"encoding/json"
	"log"
	"sync/atomic"
)

var lock int32

func init() {
	atomic.StoreInt32(&lock, 0)
}

func Send() {
	if atomic.SwapInt32(&lock, 1) == 1 {
		return
	}
	defer atomic.StoreInt32(&lock, 0)

	var message customCache.Message
	messages, err := message.ReadAll(context.Background(), customCache.Db, "messages")
	if err != nil {
		log.Printf("Error reading messages: %v", err)
		return
	}

	for _, element := range messages {
		jsonData, err := json.Marshal(element)
		if err != nil {
			log.Printf("Error marshaling message: %v", err)
			continue
		}

		err = SendData(jsonData)
		if err != nil {
			log.Printf("Error sending message: %v", err)
			continue
		}

		err = element.UpdateReceived(context.Background(), customCache.Db, "messages", element.Sequence)
		if err != nil {
			log.Printf("Error updating received status: %v", err)
		}
	}
}
