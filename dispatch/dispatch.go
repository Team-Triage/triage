package dispatch

import (
	"fmt"
	"triage/channels/toDispatch"
	"triage/channels/toFilter"
	"triage/dispatch/grpcClient/grpc"
	"triage/dispatch/grpcClient/pb"
	"triage/types"
)

func Dispatch() {
	// for {
	// networkAddress := newConsumersChan.GetMessage()
	networkAddress := "localhost:9001"
	client := grpc.ConnectToServer(networkAddress)
	go senderRoutine(client) // should also accept killchannel and networkAddress, the latter as a unique identifier for killchannel messages
	// }
	fmt.Scanln()
}

func senderRoutine(client pb.MessageHandlerClient) {
	for {
		event := toDispatch.GetMessage()
		fmt.Println("Gonna send an event!", string(event.Value))
		status := grpc.SendMessage(client, string(event.Value))

		// temporary clause because status is currently a string/should change to int or probably Bool

		var ack *types.Acknowledgment = &types.Acknowledgment{Status: int(status), Offset: int(event.TopicPartition.Offset)}

		if status < 0 { // if 'nack', add raw message to Acknowledgment struct
			ack.Event = event
		}
		toFilter.AppendMessage(ack)
	}
}
