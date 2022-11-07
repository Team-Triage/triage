package tmp

import (
	"fmt"
	"time"
	"triage/channels/acknowledgements"
	"triage/channels/messages"
	"triage/types"
)

func DummyDispatch() {
	for {
		msg := messages.GetMessage()
		fmt.Println("DummyDispatch received message")
		// make grpc call
		time.Sleep(time.Millisecond * 5) // delay to simulate grpc call
		var ack *types.Acknowledgement = &types.Acknowledgement{Offset: int(msg.TopicPartition.Offset), Status: 1, Event: msg}
		fmt.Printf("DummyDispatch sending message at offset %v to filter\n", ack.Offset)
		acknowledgements.AppendMessage(ack)
	}
}
