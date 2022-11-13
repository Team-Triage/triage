package dispatch

import (
	"fmt"

	"github.com/team-triage/triage/channels/acknowledgements"
	"github.com/team-triage/triage/channels/messages"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/team-triage/triage/channels/newConsumers"
	grpcUtils "github.com/team-triage/triage/dispatch/grpcClient/grpc"
	"github.com/team-triage/triage/dispatch/grpcClient/pb"
	"github.com/team-triage/triage/types"
)

func Dispatch() {
	for {
		networkAddress := newConsumers.GetMessage()
		fmt.Printf("DISPATCH: network address found: %v\n", networkAddress)
		client, conn := grpcUtils.MakeClient(networkAddress)
		fmt.Printf("Starting sender routine for consumer at: %v\n", networkAddress)
		go senderRoutine(client, networkAddress, conn) // should also accept killchannel and networkAddress, the latter as a unique identifier for killchannel messages
	}
}

func senderRoutine(client pb.MessageHandlerClient, networkAddress string, conn *grpc.ClientConn) {
	defer conn.Close()
	for {
		event := messages.GetMessage()
		fmt.Printf("DISPATCH: Sending event at offset %v: %v\n", int(event.TopicPartition.Offset), string(event.Value))


		respStatus, err := grpcUtils.SendMessage(client, string(event.Value))

		if err != nil {
			if status.Code(err) == codes.DeadlineExceeded {
				nack := &types.Acknowledgement{Status: -1, Offset: int(event.TopicPartition.Offset), Event: event}
				acknowledgements.AppendMessage(nack)
				fmt.Printf("SENDER ROUTINE: Deadline exceeded for offset: %v consumer: %v NACKING AND MOVING ON\n", nack.Offset, networkAddress)
				continue
			} else if status.Code(err) == codes.Unavailable {
				fmt.Printf("SENDER ROUTINE: CONSUMER DEATH at %v DETECTED - APPENDING TO MESSAGES\n", networkAddress)
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
