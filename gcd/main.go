package main

import (
	"context"
	pb "github.com/alexdnn11/go-grpc-k8s/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type server struct{}

func (s *server) Compute(ctx context.Context, r *pb.GCDRequest) (*pb.GCDResponse, error) {
	a, b := r.A, r.B
	for b != 0 {
		a, b = b, a%b
	}
	return &pb.GCDResponse{Result: a}, nil
}

func main() {

	var PORT_GRPC string
	if PORT_GRPC = os.Getenv("PORT_GRPC"); PORT_GRPC == "" {
		PORT_GRPC = "3000"
	}

	lis, err := net.Listen("tcp", PORT_GRPC)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGCDServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
