package dispatch

import (
	"fmt"
	"triage/channels/toDispatch"
	"triage/channels/toFilter"
	"triage/dispatch/grpcClient/client"
	"triage/dispatch/grpcClient/pb"
	"triage/types"
)

func Dispatch() {
	grpcClient := client.ConnectToServer("localhost:9001")
	go sendMessage()
	fmt.Scanln()
}

func sendMesssage(grpcClient *pb.MessageHandlerClient) {
	for {
		msg := toDispatch.GetMessage()
		response, err := client.SendMessage(grpcClient, msg.Value)
		if err != nil {
			fmt.Println("RUH ROH SHAGGY!")
		}
		var ack *types.Acknowledgment
		if response.Status == "nack" {
			ack = &types.Acknowledgment{Status: response.status, Offset: int(msg.TopicPartition.Offset), Event: msg}
		} else {
			ack = &types.Acknowledgment{Status: response.status, Offset: int(msg.TopicPartition.Offset)}
		}
		toFilter.AppendMessage(ack)
	}

}
