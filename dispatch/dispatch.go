package dispatch

import (
	"fmt"
	"triage/channels/acknowledgements"
	"triage/channels/messages"
	newConsumersChan "triage/channels/newConsumers"
	"triage/dispatch/grpcClient/grpc"
	"triage/dispatch/grpcClient/pb"
	"triage/types"
)

func Dispatch() {
	for {
		networkAddress := newConsumersChan.GetMessage()
		client := grpc.ConnectToServer(networkAddress)
		go senderRoutine(client) // should also accept killchannel and networkAddress, the latter as a unique identifier for killchannel messages
	}
}

func senderRoutine(client pb.MessageHandlerClient) {
	for {
		event := messages.GetMessage()
		fmt.Println("Gonna send an event!", string(event.Value))
		status := grpc.SendMessage(client, string(event.Value))

		var ack *types.Acknowledgement = &types.Acknowledgement{Status: int(status), Offset: int(event.TopicPartition.Offset)}

		if status < 0 { // if 'nack', add raw message to Acknowledgment struct
			ack.Event = event
		}
		acknowledgements.AppendMessage(ack)
	}
}
