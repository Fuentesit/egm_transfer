package main

import (
	"log"
	"net"

	//"github.com/tutorialedge/go-grpc-tutorial/chat"
	"EGM_TRANSFER/chat"
	"google.golang.org/grpc"
	grpc "google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localspot :9000")
	if err != nil {
		log.Fatalf("Failed to listen on port 9000 : %v", err)
	}
	grpcServer := grpc.NewServer()

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve grpc server over port 9000: %v", err)
	}
	
	grpc :=grpc.NewServer

	chat.RegisterChatServiceServer(grpc, )


}

