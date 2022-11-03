package client

import (
	"context"
	"log"
	"time"


	pb "triage/dispatch/grpcClient/pb" // import protobuf module
	"google.golang.org/grpc"
)

// const (
// 	address = "localhost:9001"
// )

// func messageFromChannel(c chan string) string {
// 	go channel.Pinger(channel.C)
// 	msg := channel.Printer(channel.C)
// 	return msg
// }

func ConnectToServer(address string) pb.MessageHandlerClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	
	client := pb.NewMessageHandlerClient(conn) // init client

	return client 
}

func SendMessage(client pb.MessageHandlerClient, msg string) { // will update parameter from string to proper struct
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	
	resp, err := client.SendMessage(ctx, &pb.Message{Body: msg })//messageFromChannel(channel.C)})
	
	if err != nil {
		log.Fatalf("could not get message: %v", err)
	}

	log.Printf(`Message: %s Status: %v`, resp.GetBody(), resp.GetStatus())
}
