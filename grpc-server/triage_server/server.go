package main

import (
	"context"
	"log"
	"net"
	// "status"
	// "codes"

	pb "example.com/triage-grpc/triage_main" // import protobuf module
	"google.golang.org/grpc"
	// "google.golang.org/grpc/status"
)

const (
	port = ":9001"
)

type MessageHandlerServer struct { // to register this type with grpc, embed unimpl. server inside the type
	pb.UnimplementedMessageHandlerServer
}

func (s *MessageHandlerServer) GetMessage(ctx context.Context, in *pb.Message) (*pb.MessageResponse, error) {
	// we could have msghandlerserver accept a callback 
	// a function could be called to determine code
	// return a status code.

	// eg/ positive number = ack, negative number = nack
	// try catch or if/error block
	// we send them a msg; it's on server to say what's an ack or nack
	// we'll act based on that.

	// extract into package
	// one package that's a server
	// one package that's a client
// import this package 

//have a client and server package separately. 
// for each directory, they should have a main.go file to import a server.

// is there a way in Go's pkg management to import ONLY PART OF A REPO?

// set up a dependency so that there's a 3rd repo from where client and server can get pb files and dependencies.

	// 	if err != nil {
//     errStatus := status.Convert(err)
//     log.Printf("SayHello return error: code: %d, msg: %s\n", errStatus.Code(), errStatus.Message())
// }

// 	s, ok := status.New()
	statusCode := "available"
	log.Printf("Received: %v, Status: %s", in.GetBody(), statusCode)
	return &pb.MessageResponse{Body: in.GetBody(), Status: statusCode}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer() // init new server
	pb.RegisterMessageHandlerServer(s, &MessageHandlerServer{}) // register server as a new gRPC service!
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	// s, ok := status.FromError(err)
	// fmt.Println(s, ok)
}