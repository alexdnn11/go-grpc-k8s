package main

import (
	"encoding/json"
	"fmt"
	"github.com/alexdnn11/go-grpc-k8s/pb"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"net/http"
	"os"
)

type ProofKey struct {
	ID string `json:"id"`
}

type ProofValue struct {
	Signature           []byte                   `json:"signature"`
	DataForVerification ProofDataForVerification `json:"dataForVerification"`
	State               int                      `json:"state"`
	ConsignorName       string                   `json:"consignorName"`
	Owner               string                   `json:"owner"`
	Timestamp           int64                    `json:"timestamp"`
	ShipmentID          string                   `json:"shipmentID"`
	UpdatedDate         int64                    `json:"updatedDate"`
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
	certFile = "certs/server.crt"
)

func main() {
	var PORT_GRPC, PORT_API, GCD_SERVICE_NAME string

	if PORT_GRPC = os.Getenv("PORT_GRPC"); PORT_GRPC == "" {
		PORT_GRPC = "3000"
	}
	log.Info(fmt.Sprintf("Service port: %s", PORT_GRPC))

	if PORT_API = os.Getenv("PORT_API"); PORT_API == "" {
		PORT_GRPC = "8080"
	}
	log.Info(fmt.Sprintf("Client port: %s", PORT_API))

	if GCD_SERVICE_NAME = os.Getenv("GCD_SERVICE_NAME"); GCD_SERVICE_NAME == "" {
		GCD_SERVICE_NAME = "localhost"
	}
	log.Info(fmt.Sprintf("Service name: %s", GCD_SERVICE_NAME))

	creds, err := credentials.NewClientTLSFromFile(certFile, "")
	if err != nil {
		log.Fatal("Failed to load certs: %v", err)
	}

	conn, err := grpc.Dial(GCD_SERVICE_NAME+":"+PORT_GRPC, grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatal("Dial failed: %v", err)
	}
	defer conn.Close()

	serviceClient := pb.NewServiceClient(conn)

	r := gin.Default()
	r.POST("/generate", func(ctx *gin.Context) {
		var attributes struct {
			Attributes []AttributeData `json:"attributes"`
		}
		err := ctx.BindJSON(&attributes)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		attributesBytes, err := json.Marshal(attributes.Attributes)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}

		// Call GCD service
		req := &pb.GenerateRequest{Attributes: attributesBytes}
		if res, err := serviceClient.Generate(ctx, req); err == nil {
			var proof Proof
			err := json.Unmarshal(res.Result, &proof)
			if err != nil {
				message := fmt.Sprintf("Input json is invalid. Error \"%s\"", err.Error())
				fmt.Println(message)
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			ctx.JSON(http.StatusOK, gin.H{
				"result": proof,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	r.POST("/verify", func(ctx *gin.Context) {

		var proof struct {
			Proof Proof `json:"proof"`
		}
		err := ctx.BindJSON(&proof)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		proofBytes, err := json.Marshal(proof.Proof)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		// Call GCD service
		req := &pb.VerifyRequest{Proof: proofBytes}
		if res, err := serviceClient.Verify(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": res.Result,
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	if err := r.Run(":" + PORT_API); err != nil {
		panic("Cannot serve!")
	}
}
