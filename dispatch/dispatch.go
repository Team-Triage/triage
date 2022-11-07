package dispatch

import (
	"fmt"

	"github.com/team-triage/triage/channels/acknowledgements"
	"github.com/team-triage/triage/channels/messages"

	"github.com/team-triage/triage/channels/newConsumers"
	"github.com/team-triage/triage/dispatch/grpcClient/grpc"
	"github.com/team-triage/triage/dispatch/grpcClient/pb"
	"github.com/team-triage/triage/types"
)

func Dispatch() {
	for {
		networkAddress := newConsumers.GetMessage()
		client := grpc.ConnectToServer(networkAddress)
		go senderRoutine(client) // should also accept killchannel and networkAddress, the latter as a unique identifier for killchannel messages
	}
}

func senderRoutine(client pb.MessageHandlerClient) {
	for {
		event := messages.GetMessage()
		fmt.Printf("Gonna send an event at offset %v: %v\n", int(event.TopicPartition.Offset), string(event.Value))
		status := grpc.SendMessage(client, string(event.Value))

		var ack *types.Acknowledgement = &types.Acknowledgement{Status: int(status), Offset: int(event.TopicPartition.Offset)}

		if status < 0 { // if 'nack', add raw message to Acknowledgment struct
			ack.Event = event
		}
		acknowledgements.AppendMessage(ack)
	}
}
