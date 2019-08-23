package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/hyperledger/fabric-amcl/amcl/FP256BN"
	"github.com/hyperledger/fabric/idemix"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"math/rand"
	"net"
	"os"
	"go-grpc-k8s/pb"
)

type server struct{}

type ProofKey struct {
	ID string `json:"id"`
}

type ProofValue struct {
	Signature           []byte `json:"signature"`
	DataForVerification []byte `json:"dataForVerification"`
	State               int    `json:"state"`
	ConsignorName       string `json:"consignorName"`
	Owner               string `json:"owner"`
	Timestamp           int64  `json:"timestamp"`
	ShipmentID          string `json:"shipmentID"`
	UpdatedDate         int64  `json:"updatedDate"`
}

type ProofDataForVerification struct {
	Disclosure          []byte   `json:"disclosure"`
	Ipk                 []byte   `json:"ipk"`
	Msg                 []byte   `json:"msg"`
	AttributeValuesHash [][]byte `json:"attributeValuesHash"`
	AttributeValues     []string `json:"attributeValues"`
	RhIndex             int      `json:"rhIndex"`
	RevPk               string   `json:"revPk"`
	Epoch               int      `json:"epoch"`
}

type Proof struct {
	Key   ProofKey   `json:"key"`
	Value ProofValue `json:"value"`
}

type AttributeData struct {
	AttributeName       string `json:"attributeName"`
	AttributeValue      string `json:"attributeValue"`
	AttributeDisclosure int    `json:"attributeDisclosure"`
}

const (
	certFile = "certs/server/server.crt"
	keyFile  = "certs/server/server.key"
)

func (s *server) Generate(ctx context.Context, r *pb.GenerateRequest) (*pb.GenerateResponse, error) {

	log.Info(fmt.Sprintf("### Generate started ###"))
	log.Debug(ctx)

	proof := Proof{}
	data := ProofDataForVerification{}

	// making arrays of attributes names and values
	rng, err := idemix.GetRand()

	var attributesArray []AttributeData
	err = json.Unmarshal([]byte(r.Attributes), &attributesArray)
	if err != nil {
		message := fmt.Sprintf("Input json is invalid. Error \"%s\"", err.Error())
		fmt.Println(message)
		return &pb.GenerateResponse{Result: nil}, err
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
		disclosure[i] = byte(attributesArray[i].AttributeDisclosure)
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
		return &pb.GenerateResponse{Result: nil}, errors.New(message)
	}

	// create a new key pair
	key, err := idemix.NewIssuerKey(AttributeNames, rng)
	if err != nil {
		message := fmt.Sprintf("Issuer key generation should have succeeded but gave error \"%s\"", err.Error())
		fmt.Println(message)
		return &pb.GenerateResponse{Result: nil}, errors.New(message)
	}

	// check that the key is valid
	err = key.GetIpk().Check()
	if err != nil {
		message := fmt.Sprintf("Issuer public key should be valid")
		fmt.Println(message)
		return &pb.GenerateResponse{Result: nil}, errors.New(message)
	}

	// issuance
	sk := idemix.RandModOrder(rng)
	ni := idemix.RandModOrder(rng)
	m := idemix.NewCredRequest(sk, ni, key.Ipk, rng)

	cred, err := idemix.NewCredential(key, m, attrs, rng)

	// generate a revocation key pair
	revocationKey, err := idemix.GenerateLongTermRevocationKey()

	// create CRI that contains no revocation mechanism
	epoch := 0
	cri, err := idemix.CreateCRI(revocationKey, []*FP256BN.BIG{}, epoch, idemix.ALG_NO_REVOCATION, rng)
	if err != nil {
		message := fmt.Sprintf("Create CRI return error: %s", err.Error())
		fmt.Println(message)
		return &pb.GenerateResponse{Result: nil}, errors.New(message)
	}

	// signing selective disclosure
	Nym, RandNym := idemix.MakeNym(sk, key.Ipk, rng)
	sig, err := idemix.NewSignature(cred, sk, Nym, RandNym, key.Ipk, disclosure, msg, rhindex, cri, rng)
	if err != nil {
		message := fmt.Sprintf("Idemix NewSignature return error: %s", err.Error())
		fmt.Println(message)
		return &pb.GenerateResponse{Result: nil}, errors.New(message)
	}

	attributeValuesBytes := make([][]byte, len(attrs))
	for i := 0; i < len(attrs); i++ {
		row := make([]byte, FP256BN.MODBYTES)
		attributeValue := attrs[i]
		attributeValue.ToBytes(row)
		attributeValuesBytes[i] = row
	}

	sigBytes, err := json.Marshal(sig)
	if err != nil {
		message := fmt.Sprintf("Signature marshaling error: %s", err.Error())
		fmt.Println(message)
		return &pb.GenerateResponse{Result: nil}, errors.New(message)
	}

	ipkBytes, err := json.Marshal(key.Ipk)
	if err != nil {
		message := fmt.Sprintf("Ipk marshaling error: %s", err.Error())
		fmt.Println(message)
		return &pb.GenerateResponse{Result: nil}, errors.New(message)
	}

	proof.Value.Signature = sigBytes
	data.Disclosure = disclosure
	data.Ipk = ipkBytes
	data.Msg = msg
	data.AttributeValuesHash = attributeValuesBytes
	data.AttributeValues = attributeValues
	data.RhIndex = rhindex
	data.RevPk = encode(&revocationKey.PublicKey)
	data.Epoch = epoch

	dataBytes, err := json.Marshal(data)
	if err != nil {
		message := fmt.Sprintf("ProofDataForVerification marshaling error: %s", err.Error())
		fmt.Println(message)
		return &pb.GenerateResponse{Result: nil}, errors.New(message)
	}

	proof.Value.DataForVerification = dataBytes

	result, err := json.Marshal(proof)
	if err != nil {
		message := fmt.Sprintf("cannot marshal. Error \"%s\"", err.Error())
		fmt.Println(message)
		return &pb.GenerateResponse{Result: nil}, err
	}

	log.Info(fmt.Sprintf("### Generate successfully completed ###"))

	return &pb.GenerateResponse{Result: result}, nil
}

func (s *server) Verify(ctx context.Context, r *pb.VerifyRequest) (*pb.VerifyResponse, error) {

	log.Info(fmt.Sprintf("### Verify started ###"))
	log.Debug(ctx)

	proof := Proof{}
	data :=ProofDataForVerification{}

	err := json.Unmarshal([]byte(r.Proof), &proof)
	if err != nil {
		message := fmt.Sprintf("Input json is invalid. Error \"%s\"", err.Error())
		fmt.Println(message)
		return &pb.VerifyResponse{Result: false}, err
	}

	var sig *idemix.Signature
	var ipk *idemix.IssuerPublicKey

	err = json.Unmarshal([]byte(proof.Value.Signature), &sig)
	if err != nil {
		message := fmt.Sprintf("Input json is invalid. Error \"%s\"", err.Error())
		fmt.Println(message)
		return &pb.VerifyResponse{Result: false}, err
	}

	err = json.Unmarshal([]byte(proof.Value.DataForVerification), &data)
	if err != nil {
		message := fmt.Sprintf("Input json is invalid. Error \"%s\"", err.Error())
		fmt.Println(message)
		return &pb.VerifyResponse{Result: false}, err
	}

	err = json.Unmarshal([]byte(data.Ipk), &ipk)
	if err != nil {
		message := fmt.Sprintf("Input json is invalid. Error \"%s\"", err.Error())
		fmt.Println(message)
		return &pb.VerifyResponse{Result: false}, err
	}

	attributeValuesBytes := make([]*FP256BN.BIG, len(data.AttributeValuesHash))

	for i := range data.AttributeValuesHash {
		attributeValuesBytes[i] = FP256BN.FromBytes(data.AttributeValuesHash[i])
	}
	err = sig.Ver(data.Disclosure,
		ipk,
		data.Msg,
		attributeValuesBytes,
		data.RhIndex,
		decode(data.RevPk),
		data.Epoch)

	if err != nil {
		message := fmt.Sprintf("Signature verification was failed. Error: %s", err.Error())
		fmt.Println(message)
		return &pb.VerifyResponse{Result: false}, err
	}

	log.Info(fmt.Sprintf("### Verify successfully completed ###"))

	return &pb.VerifyResponse{Result: true}, nil
}

func encode(publicKey *ecdsa.PublicKey) string {
	x509EncodedPub, _ := x509.MarshalPKIXPublicKey(publicKey)
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})

	return string(pemEncodedPub)
}

func decode(pemEncodedPub string) *ecdsa.PublicKey {
	blockPub, _ := pem.Decode([]byte(pemEncodedPub))
	x509EncodedPub := blockPub.Bytes
	genericPublicKey, _ := x509.ParsePKIXPublicKey(x509EncodedPub)
	publicKey := genericPublicKey.(*ecdsa.PublicKey)

	return publicKey
}

func main() {

	var PORT_GRPC, TLS string

	if PORT_GRPC = os.Getenv("PORT_GRPC"); PORT_GRPC == "" {
		PORT_GRPC = "5000"
	}
	log.Info(fmt.Sprintf("Service port: %s", PORT_GRPC))

	if TLS = os.Getenv("TLS_ENABLE"); TLS == "" {
		TLS = "false"
	}
	log.Info(fmt.Sprintf("TLS_ENABLE: %s", TLS))

	lis, err := net.Listen("tcp", ":"+PORT_GRPC)
	if err != nil {
		log.Fatal("Failed to listen: %v", err)
	}

	creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
	if err != nil {
		log.Fatal("Failed to load certs: %v", err)
	}

	var s *grpc.Server

	if TLS == "true" {
		s = grpc.NewServer(grpc.Creds(creds))
	} else {
		s = grpc.NewServer()
	}

	pb.RegisterIdemixServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatal("Failed to serve: %v", err)
	}
}
