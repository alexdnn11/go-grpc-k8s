package main

import (
	"fmt"
	"github.com/alexdnn11/go-grpc-k8s/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {
	var PORT_GRPC, PORT_API string
	if PORT_GRPC = os.Getenv("PORT_GRPC"); PORT_GRPC == "" {
		PORT_GRPC = "3000"
	}
	if PORT_API = os.Getenv("PORT_API"); PORT_API == "" {
		PORT_GRPC = "8080"
	}

	conn, err := grpc.Dial("localhost:"+PORT_GRPC, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Dial failed: %v", err)
	}
	defer conn.Close()
	gcdClient := pb.NewGCDServiceClient(conn)

	r := gin.Default()
	fmt.Println("Api started!")
	r.GET("/gcd/:a/:b", func(ctx *gin.Context) {
		// Parse parameters
		a, err := strconv.ParseUint(ctx.Param("a"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter A"})
			return
		}
		b, err := strconv.ParseUint(ctx.Param("b"), 10, 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter B"})
			return
		}
		// Call GCD service
		req := &pb.GCDRequest{A: a, B: b}
		if res, err := gcdClient.Compute(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint(res.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	if err := r.Run(":" + PORT_API); err != nil {
		panic("Cannot serve!")
	}
}
