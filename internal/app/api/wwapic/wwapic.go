package main

import (
	"context"
	"log"
	"time"

	wwapi "github.com/hpcng/warewulf/internal/pkg/api/routes/wwapiv1"

	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"

	"google.golang.org/grpc/credentials/insecure" // TODO: Block server on startup if in insecure mode.
)

// wwapic is intended as a sample wwapi client.

func main() {
	log.Println("Client running")

	// TODO: Remove hardcoded port. Config file.
	conn, err := grpc.Dial(":9872", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := wwapi.NewWWApiClient(conn)

	request := &emptypb.Empty{}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.Version(ctx, request)
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("Version Response: %v", response)
}

