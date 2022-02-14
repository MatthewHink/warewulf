package main

import (
	"context"
	"log"
	"time"

	wwapi "github.com/hpcng/warewulf/pkg/api/wwapiv1"

	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	// wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	log.Println("Client running")

	conn, err := grpc.Dial(":9872", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Println("Error1:")
		log.Fatalln(err)
	}
	defer conn.Close()

	client := wwapi.NewWWApiClient(conn)

	request := &emptypb.Empty{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.Version(ctx, request)
	if err != nil {
		log.Println("Error2:")
		log.Fatalln(err)
	}

	log.Printf("Response: %v", response.Value)
}
