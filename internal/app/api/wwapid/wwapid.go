package main

import (
	"context"
	"log"
	"net"
	"fmt"
	"os"

	wwapi "github.com/hpcng/warewulf/internal/pkg/api/routes/wwapiv1"
	"github.com/hpcng/warewulf/internal/pkg/api/node"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"

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
	config, err := wwapidconf.New("")
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

// TODO: Needs testing.
// TODO: Need nil check on nodeNames arrays.
func (s *apiServer) NodeAdd(ctx context.Context, request *wwapi.NodeAddParameter) (response *emptypb.Empty, err error) {

	log.Println("NodeAdd start")
	log.Printf("request: %T, %#v\n", request, request)

	if request == nil {
		return response, status.Errorf(codes.InvalidArgument, "nil request")
	}
	
	response = new(emptypb.Empty)
	err = node.NodeAdd(request)
	return
}

func (s *apiServer) NodeDelete(ctx context.Context, request *wwapi.NodeDeleteParameter) (response *emptypb.Empty, err error) {

	log.Println("NodeDelete start")
	log.Printf("request: %T, %#v\n", request, request)

	if request == nil {
		return response, status.Errorf(codes.InvalidArgument, "nil request")
	}

	response = new(emptypb.Empty)
	err = node.NodeDelete(request)
	return
}

func (s *apiServer) NodeList(ctx context.Context, request *wwapi.NodeNames) (response *wwapi.NodeListResponse, err error) {

	log.Println("NodeList start")
	log.Printf("request: %T, %#v\n", request, request)

	if request == nil {
		return response, status.Errorf(codes.InvalidArgument, "nil request")
	}

	var nodes []*wwapi.NodeInfo
	nodes, err = node.NodeList(request.NodeNames)
	if err != nil {
		return
	}

	response = &wwapi.NodeListResponse{
		Nodes: nodes,
	}
	return
}

func (s *apiServer) Version(ctx context.Context, request *emptypb.Empty) (response *wwapi.VersionResponse, err error) {

	response = &wwapi.VersionResponse{
		ApiPrefix: apiPrefix,
		ApiVersion: apiVersion,
		WarewulfVersion: version.GetVersion(),
	}
	return
}
