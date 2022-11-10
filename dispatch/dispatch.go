package dispatch

import (
	"fmt"

	"github.com/team-triage/triage/channels/acknowledgements"
	"github.com/team-triage/triage/channels/messages"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/team-triage/triage/channels/newConsumers"
	"github.com/team-triage/triage/dispatch/grpcClient/grpc"
	"github.com/team-triage/triage/dispatch/grpcClient/pb"
	"github.com/team-triage/triage/types"
)

func Dispatch() {
	for {
		networkAddress := newConsumers.GetMessage()
		fmt.Println("DISPATCH: network address found!", networkAddress)
		client := grpc.MakeClient(networkAddress)
		fmt.Println("Starting sender routine")
		go senderRoutine(client) // should also accept killchannel and networkAddress, the latter as a unique identifier for killchannel messages
	}
}

func senderRoutine(client pb.MessageHandlerClient) {
	for {
		event := messages.GetMessage()
		fmt.Printf("DISPATCH: Sending event at offset %v: %v\n", int(event.TopicPartition.Offset), string(event.Value))

		fmt.Printf("DISPATCH: Sending event topic :%v\n partition: %v\n offset: %v\n key: %v\n value: %v\n timestamp: %v\n headers: %v\n",
			&event.TopicPartition.Topic,
			event.TopicPartition.Partition,
			int(event.TopicPartition.Offset),
			string(event.Key),
			string(event.Value),
			event.Timestamp,
			event.Headers,
		)

		respStatus, err := grpc.SendMessage(client, string(event.Value))

		if err != nil {
			if status.Code(err) == codes.DeadlineExceeded {
				nack := &types.Acknowledgement{Status: -1, Offset: int(event.TopicPartition.Offset), Event: event}
				acknowledgements.AppendMessage(nack)
				fmt.Println("SENDER ROUTINE: DEADLINE EXCEEDED - NACKING AND MOVING ON")
				continue
			} else if status.Code(err) == codes.Unavailable {
				fmt.Println("SENDER ROUTINE: CONSUMER DEATH DETECTED - APPENDING TO MESSAGES")
				messages.AppendMessage(event)
				break
			}
		}

		var ack *types.Acknowledgement = &types.Acknowledgement{Status: respStatus, Offset: int(event.TopicPartition.Offset)}

		if respStatus < 0 { // if 'nack', add raw message to Acknowledgment struct
			ack.Event = event
		}
		acknowledgements.AppendMessage(ack)
	}
}
