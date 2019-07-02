package main

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	pb "github.com/alexdnn11/go-grpc-k8s/pb"
	"github.com/hyperledger/fabric-amcl/amcl/FP256BN"
	"github.com/hyperledger/fabric/idemix"
	"google.golang.org/grpc"
	"log"
	"math/rand"
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

	// making arrays of attributes names and values
	rng, err := idemix.GetRand()

	var attributesArray []AttributeData
	err = json.Unmarshal([]byte(r.Attributes), &attributesArray)
	if err != nil {
		message := fmt.Sprintf("Input json is invalid. Error \"%s\"", err.Error())
		fmt.Println(message)
		return &pb.GCDResponse{Result: nil}, err
	}

	AttributeNames := make([]string, len(attributesArray))
	attrs := make([]*FP256BN.BIG, len(AttributeNames))
	disclosure := make([]byte, len(attributesArray))
	msg := make([]byte, len(attributesArray))
	attributeValues := make([]string, len(attributesArray))
	var rhindex int

	for i := range attributesArray {
		h := sha256.New()
		// make hash from value of attribute
		h.Write([]byte(attributesArray[i].AttributeValue))
		attrs[i] = FP256BN.FromBytes(h.Sum(nil))
		AttributeNames[i] = attributesArray[i].AttributeName
		disclosure[i] = attributesArray[i].AttributeDisclosure
		msg[i] = byte(i)
		if attributesArray[i].AttributeDisclosure == 0 {
			rhindex = i
			// fill hidden field random value
			attrs[i] = FP256BN.NewBIGint(rand.Intn(10000))
		} else {
			attributeValues[i] = attributesArray[i].AttributeValue
		}
	}

	// check Disclosure[rhIndex] == 0
	if attributesArray[rhindex].AttributeDisclosure != 0 {
		message := fmt.Sprintf("Idemix requires the revocation handle to remain undisclosed (i.e., Disclosure[rhIndex] == 0). But we have \"%d\"", attributesArray[rhindex].AttributeDisclosure)
		fmt.Println(message)
		return &pb.GCDResponse{Result: nil}, errors.New(message)
	}

	// create a new key pair
	key, err := idemix.NewIssuerKey(AttributeNames, rng)
	if err != nil {
		message := fmt.Sprintf("Issuer key generation should have succeeded but gave error \"%s\"", err.Error())
		fmt.Println(message)
		return &pb.GCDResponse{Result: nil}, errors.New(message)
	}

	// check that the key is valid
	err = key.GetIpk().Check()
	if err != nil {
		message := fmt.Sprintf("Issuer public key should be valid")
		fmt.Println(message)
		return &pb.GCDResponse{Result: nil}, errors.New(message)
	}

	// issuance
	sk := idemix.RandModOrder(rng)
	ni := idemix.RandModOrder(rng)
	m := idemix.NewCredRequest(sk, idemix.BigToBytes(ni), key.Ipk, rng)

	cred, err := idemix.NewCredential(key, m, attrs, rng)

	// generate a revocation key pair
	revocationKey, err := idemix.GenerateLongTermRevocationKey()

	// create CRI that contains no revocation mechanism
	epoch := 0
	cri, err := idemix.CreateCRI(revocationKey, []*FP256BN.BIG{}, epoch, idemix.ALG_NO_REVOCATION, rng)
	if err != nil {
		message := fmt.Sprintf("Create CRI return error: %s", err.Error())
		fmt.Println(message)
		return &pb.GCDResponse{Result: nil}, errors.New(message)
	}

	// signing selective disclosure
	Nym, RandNym := idemix.MakeNym(sk, key.Ipk, rng)
	sig, err := idemix.NewSignature(cred, sk, Nym, RandNym, key.Ipk, disclosure, msg, rhindex, cri, rng)
	if err != nil {
		message := fmt.Sprintf("Idemix NewSignature return error: %s", err.Error())
		fmt.Println(message)
		return &pb.GCDResponse{Result: nil}, errors.New(message)
	}

	attributeValuesBytes := make([][]byte, len(attrs))
	for i := 0; i < len(attrs); i++ {
		row := make([]byte, FP256BN.MODBYTES)
		attributeValue := attrs[i]
		attributeValue.ToBytes(row)
		attributeValuesBytes[i] = row
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
