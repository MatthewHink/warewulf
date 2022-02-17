package main

import (
	"context"
	"log"
	"net"
	"fmt"
	"os"

	wwapi "github.com/hpcng/warewulf/internal/pkg/api/routes/wwapiv1"

	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	//wrapperspb "google.golang.org/protobuf/types/known/wrapperspb""
	wwapidconf "github.com/hpcng/warewulf/internal/pkg/wwapidconf"
)

type apiServer struct {
	wwapi.UnimplementedWWApiServer
}

func main() {
	log.Println("Server running")

	// Read the config file.
	config, err := wwapidconf.New()
	if err != nil {
		// TODO: log
		fmt.Printf("err: %v", err)
		os.Exit(1)
	}
	fmt.Printf("config: %#v\n", config)


	listen, err := net.Listen("tcp", ":9872") // TODO: Port in config file.
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
	wwapi.RegisterWWApiServer(grpcServer, &apiServer{})

	log.Fatalln(grpcServer.Serve(listen))
}

// Api implementation.
func (s *apiServer) Version(ctx context.Context, request *emptypb.Empty) (response *wwapi.VersionResponse, err error) {
	//str = wrapperspb.String("Version 0.0.0")
	// TODO: Fix hardcoding.
	// Version in config file and Makefile.
	// Warewulf version from warewulf.
	response = &wwapi.VersionResponse{
		ApiPrefix: "v1",
		ApiVersion: "1.0.0",
		WarewulfVersion: "4.3",
	}
	return
}