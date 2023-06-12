package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"workspace"

	"google.golang.org/grpc"
)

type Server struct {
	workspace.UnimplementedTextProcessorServer
}

func (s *Server) Process(ctx context.Context, processRequest *workspace.ProcessRequest) (*workspace.ProcessResponse, error) {
	log.Printf("[gRPC] Receive message body from client: %+v", processRequest)
	return &workspace.ProcessResponse{
		Words:   []string{"test"},
		Message: "test",
	}, nil
}

func main() {
	fmt.Println("[gRPC] Starting server")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 9000))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := Server{}

	grpcServer := grpc.NewServer()

	workspace.RegisterTextProcessorServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
