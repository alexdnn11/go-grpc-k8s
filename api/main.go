package main

import (
	"encoding/json"
	"fmt"
	"github.com/alexdnn11/go-grpc-k8s/pb"
	"github.com/gin-gonic/gin"
	"github.com/hyperledger/fabric/idemix"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
)

type ProofKey struct {
	ID string `json:"id"`
}

type ProofValue struct {
	SnapShot            *idemix.Signature        `json:"snapShot"`
	DataForVerification ProofDataForVerification `json:"dataForVerification"`
	State               int                      `json:"state"`
	ConsignorName       string                   `json:"consignorName"`
	Owner               string                   `json:"owner"`
	Timestamp           int64                    `json:"timestamp"`
	ShipmentID          string                   `json:"shipmentID"`
	UpdatedDate         int64                    `json:"updatedDate"`
}

type ProofDataForVerification struct {
	Disclosure          []byte                  `json:"disclosure"`
	Ipk                 *idemix.IssuerPublicKey `json:"ipk"`
	Msg                 []byte                  `json:"msg"`
	AttributeValuesHash [][]byte                `json:"attributeValuesHash"`
	AttributeValues     []string                `json:"attributeValues"`
	RhIndex             int                     `json:"rhIndex"`
	RevPk               string                  `json:"revPk"`
	Epoch               int                     `json:"epoch"`
}

type Proof struct {
	Key   ProofKey   `json:"key"`
	Value ProofValue `json:"value"`
}

type AttributeData struct {
	AttributeName       string `json:"attributeName"`
	AttributeValue      string `json:"attributeValue"`
	AttributeDisclosure byte   `json:"attributeDisclosure"`
}

func main() {
	var PORT_GRPC, PORT_API, GCD_SERVICE_NAME string
	if PORT_GRPC = os.Getenv("PORT_GRPC"); PORT_GRPC == "" {
		PORT_GRPC = "3000"
	}
	if PORT_API = os.Getenv("PORT_API"); PORT_API == "" {
		PORT_GRPC = "8080"
	}
	if GCD_SERVICE_NAME = os.Getenv("GCD_SERVICE_NAME"); GCD_SERVICE_NAME == "" {
		GCD_SERVICE_NAME = "localhost"
	}

	conn, err := grpc.Dial(GCD_SERVICE_NAME+":"+PORT_GRPC, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	defer conn.Close()
	gcdClient := pb.NewGCDServiceClient(conn)

	r := gin.Default()
	fmt.Println("Api started!")
	r.POST("/compute", func(ctx *gin.Context) {
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
		req := &pb.GCDRequest{Attributes: attributesBytes}
		if res, err := gcdClient.Compute(ctx, req); err == nil {
			var proof []Proof
			err := json.Unmarshal(res.Result, &proof)
			if err != nil {
				message := fmt.Sprintf("Input json is invalid. Error \"%s\"", err.Error())
				fmt.Println(message)
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			}
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(proof),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	if err := r.Run(":" + PORT_API); err != nil {
		panic("Cannot serve!")
	}
}
