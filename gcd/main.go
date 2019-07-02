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

	var attributesArray []AttributeData
	err := json.Unmarshal([]byte(r.Attributes), &attributesArray)
	if err != nil {
		message := fmt.Sprintf("Input json is invalid. Error \"%s\"", err.Error())
		fmt.Println(message)
		return &pb.GCDResponse{Result: nil}, err
	}
	for i, _ := range attributesArray {
		attributesArray[i].AttributeName = "ECHO"
	}

	result, err := json.Marshal(attributesArray)
	if err != nil {
		message := fmt.Sprintf("cannot marshal. Error \"%s\"", err.Error())
		fmt.Println(message)
		return &pb.GCDResponse{Result: nil}, err
	}

	return &pb.GCDResponse{Result: result}, nil
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
