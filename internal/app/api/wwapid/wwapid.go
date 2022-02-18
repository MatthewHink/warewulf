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
	"github.com/hpcng/warewulf/internal/pkg/version"
)

type apiServer struct {
	wwapi.UnimplementedWWApiServer
}

var apiPrefix string
var apiVersion string

func main() {
	log.Println("Server running")

	// Read the config file.
	config, err := wwapidconf.New("") // TODO: path parameter
	if err != nil {
		// TODO: wwapi log
		log.Printf("err: %v", err)
		os.Exit(1)
	}
	log.Printf("config: %#v\n", config) // TODO: log

	// Pull out config variables and log.
	apiPrefix = config.ApiPrefix
	apiVersion = config.ApiVersion
	servicePort := config.Port
	log.Printf("Starting wwapid. Version %s. Port %d. ApiPrefix %s\n",
		apiVersion, servicePort, apiPrefix)
	portString := fmt.Sprintf(":%d", servicePort)

	listen, err := net.Listen("tcp", portString)
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
	wwapi.RegisterWWApiServer(grpcServer, &apiServer{})

	log.Fatalln(grpcServer.Serve(listen))
}

// Api implementation.

func (s *apiServer) Version(ctx context.Context, request *emptypb.Empty) (response *wwapi.VersionResponse, err error) {

	response = &wwapi.VersionResponse{
		ApiPrefix: apiPrefix,
		ApiVersion: apiVersion,
		WarewulfVersion: version.GetVersion(),
	}
	return
}