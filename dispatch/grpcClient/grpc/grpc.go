package grpc

import (
	"context"
	"fmt"
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

	// defer conn.Close()

	client := pb.NewMessageHandlerClient(conn) // init client

	return client
}

func SendMessage(client pb.MessageHandlerClient, msg string) int32 { // will update parameter from string to proper struct
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	fmt.Println("GRPC is about to send a message!", msg)
	resp, err := client.SendMessage(ctx, &pb.Message{Body: msg})
	fmt.Println(resp)
	if err != nil {
		log.Fatalf("could not get message: %v", err)
	}
	return resp.GetStatus() // "ack" or "nack"
}
