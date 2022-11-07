package tmp

import (
	"fmt"
	"time"

	"github.com/team-triage/triage/channels/acknowledgements"
	"github.com/team-triage/triage/channels/messages"
	"github.com/team-triage/triage/types"
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
