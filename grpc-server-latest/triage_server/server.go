package server

import (
	"context"
	"log"
	"net"
	// "fmt"
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

// func unaryInterceptorImplementation(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
// 	fmt.Print(resp)
// 	md, err := handler(ctx, req)
// 	if err != nil {
// 		log.Fatalf("error in grpc service %v", err)
// 	}
// 	return md, err
// }



func messageProcessor(msg string) int {
	result := 0
	if msg == "ping" {
		result = 1
	} else {
		result = -1
	}
	return result
}

// var functionStore func();
var functionStore func(msg string) int
// init var without assigning
// define onMessage function
func OnMessage(f func(msg string) int) {
	functionStore = f
}
// func onMessage(function ) {
// 	functionStore = function;
// }


func (s *MessageHandlerServer) GetMessage(ctx context.Context, in *pb.Message) (*pb.MessageResponse, error) {
	
	// statusInt := messageProcessor(in.GetBody())
	statusInt := functionStore(in.GetBody()) 
	statusCode := "no status" 
	if statusInt > 0 {
		statusCode = "ack" 
		} else {
			statusCode = "nack" 
		}// this should be a result of processing function
	log.Printf("Received: %v, Status: %s", in.GetBody(), statusCode)
	// on message function handles message here and saves it in a variable
	// responses channel
	return &pb.MessageResponse{Body: in.GetBody(), Status: statusCode}, nil
	// func onMessage -> pass us a fxn that is saved in a var
}
// the interceptor has to call the processing function with the incoming message.
// the result of that processing function should determine whether it's "ack" or "nack"
// if it's "ack" that goes to filter, otherwise it goes to reaper

func StartServer() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}



	s := grpc.NewServer()//(grpc.UnaryInterceptor(unaryInterceptorImplementation)) // init new server
	pb.RegisterMessageHandlerServer(s, &MessageHandlerServer{}) // register server as a new gRPC service!
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	// s, ok := status.FromError(err)
	// fmt.Println(s, ok)
}

/*
comment history:
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
// call function from 57 here and pass it the msg -> returns a # to represent an ack or nack
*/