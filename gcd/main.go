package main

import (
	"context"
	"encoding/json"
	"fmt"
	pb "github.com/alexdnn11/go-grpc-k8s/pb"
	"google.golang.org/grpc"
	"log"
	"net"
	"os"
)

type server struct{}

type AttributeData struct {
	AttributeName       string `json:"attributeName"`
	AttributeValue      string `json:"attributeValue"`
	AttributeDisclosure byte   `json:"attributeDisclosure"`
}

func (s *server) Compute(ctx context.Context, r *pb.GCDRequest) (*pb.GCDResponse, error) {
	//a, b := r.A, r.B
	//for b != 0 {
	//	a, b = b, a%b
	//}
	var attributesArray []AttributeData
	err := json.Unmarshal([]byte(r.attributes), &attributesArray)
	if err != nil {
		message := fmt.Sprintf("Input json is invalid. Error \"%s\"", err.Error())
		fmt.Println(message)
		return &pb.GCDResponse{Result: []byte{""}}, err
	}

	return &pb.GCDResponse{Result: a}, nil
}

func main() {

	var PORT_GRPC string
	if PORT_GRPC = os.Getenv("PORT_GRPC"); PORT_GRPC == "" {
		PORT_GRPC = "3000"
	}
	fmt.Println(PORT_GRPC)

	lis, err := net.Listen("tcp", ":"+PORT_GRPC)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGCDServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
