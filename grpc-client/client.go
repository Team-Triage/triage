package main

import (
	"context"
	"log"
	"time"


	pb "example.com/triage-grpc/triage_main" // import protobuf module
	"google.golang.org/grpc"
	"example.com/triage-grpc/channel"
)

const (
	address = "localhost:9001"
)

func messageFromChannel(c chan string) string {
	go channel.Pinger(channel.C)
	msg := channel.Printer(channel.C)
	return msg
}

func main() {
	
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()
	c := pb.NewMessageHandlerClient(conn) // init client

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// var messageOne = "TEAM 3 LET'S GO"
	
	resp, err := c.GetMessage(ctx, &pb.Message{Body: messageFromChannel(channel.C)})
	if err != nil {
		log.Fatalf("could not get message: %v", err)
	}
	log.Printf(`Message: %s Status: %v`, resp.GetBody(), resp.GetStatus())
	// if resp.GetStatus() === "ack" , send to Filter
	// else send to Reaper
}