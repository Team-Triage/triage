package dispatch

import (
	newConsumersChan "triage/channels/newConsumers"
	"triage/channels/toDispatch"
	"triage/channels/toFilter"
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
		event := toDispatch.GetMessage()
		status := grpc.SendMessage(client, string(event.Value))

		var statusInt int
		// temporary clause because status is currently a string/should change to int or probably Bool
		if status == "nack" {
			statusInt = -1
		} else {
			statusInt = 1
		}
		var ack *types.Acknowledgment = &types.Acknowledgment{Status: statusInt, Offset: int(event.TopicPartition.Offset)}

		if statusInt < 0 { // if 'nack', add raw message to Acknowledgment struct
			ack.Event = event
		}
		toFilter.AppendMessage(ack)
	}
}
