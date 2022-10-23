package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"os"

	"google.golang.org/grpc"
	pb "github.com/phucvin/project-teneng/services/servicerouter/servicerouter/proto"
)

// server is used to implement servicerouter.ServiceRouterServer.
type server struct {
	pb.UnimplementedServiceRouterServer
}

// Invoke implements servicerouter.ServiceRouterServer
func (s *server) Invoke(ctx context.Context, in *pb.InvokeRequest) (*pb.InvokeResponse, error) {
	log.Printf("Received: %v", in.GetDescription())
	return &pb.InvokeResponse{Description: "Received " + in.GetDescription()}, nil
}

func main() {
    port, err := strconv.Atoi(os.Getenv("PORT"))
    if err != nil {
		port = 50051
    }

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterServiceRouterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}