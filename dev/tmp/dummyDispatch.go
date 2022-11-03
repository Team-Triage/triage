package tmp

import (
	"fmt"
	"time"
	"triage/channels/toDispatch"
	"triage/channels/toFilter"
	"triage/types"
)

func DummyDispatch() {
	for {
		msg := toDispatch.GetMessage()
		fmt.Println("DummyDispatch received message")
		// make grpc call
		time.Sleep(time.Millisecond * 5) // delay to simulate grpc call
		var ack *types.Acknowledgment = &types.Acknowledgment{Offset: int(msg.TopicPartition.Offset), Status: 1, Event: msg}
		fmt.Printf("DummyDispatch sending message at offset %v to filter\n", ack.Offset)
		toFilter.AppendMessage(ack)
	}
}
