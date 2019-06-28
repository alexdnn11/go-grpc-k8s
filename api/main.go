package main

import (
	"fmt"
	"github.com/alexdnn11/go-grpc-k8s/pb"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strconv"
)

const (
	grpcAddress = "localhost:3000"
	hostAddress = "localhost:8080"
)

func main() {
	conn, err := grpc.Dial(grpcAddress, grpc.WithInsecure())
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

	if err := r.Run(hostAddress); err != nil {
		panic("Cannot serve!")
	}
}
