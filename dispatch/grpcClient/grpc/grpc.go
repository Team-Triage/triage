package grpc

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/team-triage/triage/dispatch/grpcClient/pb" // import protobuf module

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func ConnectToServer(address string) pb.MessageHandlerClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:    time.Second * 5, // how long we wait to hear back from the server before closing connection
		Timeout: time.Second * 1, // frequency of pings
	}))

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	// defer conn.Close()

	client := pb.NewMessageHandlerClient(conn) // init client

	return client
}

func SendMessage(client pb.MessageHandlerClient, msg string) int32 { // will update parameter from string to proper struct
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	fmt.Println("GRPC CLIENT: Sending message!", msg)
	resp, err := client.SendMessage(ctx, &pb.Message{Body: msg})
	fmt.Println(resp)
	if err != nil {
		log.Fatalf("could not get message: %v", err)
	}
	return resp.GetStatus() // "ack" or "nack"
}
