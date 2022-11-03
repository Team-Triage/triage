package dispatch

import (
	"fmt"
	"triage/channels/toDispatch"
	"triage/channels/toFilter"
	"triage/dispatch/grpcClient/client"
	"triage/dispatch/grpcClient/pb"
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
		
		toFilter.AppendMessage(msg)
	}

}