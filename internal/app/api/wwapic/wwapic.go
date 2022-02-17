package main

import (
	"context"
	"log"
	"time"

	wwapi "github.com/hpcng/warewulf/internal/pkg/api/routes/wwapiv1"

	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	// wrapperspb "google.golang.org/protobuf/types/known/wrapperspb"

	"google.golang.org/grpc/credentials/insecure" // TODO: Remove
)

// wwapic is intended as a sample wwapi client.

func main() {
	log.Println("Client running")

	// TODO: Remove hardcoded port. Config file.
	//conn, err := grpc.Dial(":9872", grpc.WithInsecure(), grpc.WithBlock())
	conn, err := grpc.Dial(":9872", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
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

	log.Printf("Response: %v", response)
}

