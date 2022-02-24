package main

import (
	"context"
	"log"
	"net"
	"fmt"
	"os"

	wwapi "github.com/hpcng/warewulf/internal/pkg/api/routes/wwapiv1"
	"github.com/hpcng/warewulf/internal/pkg/api/container"
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

// ContainerDelete deletes one or more containers from Warewulf.
func (s *apiServer) ContainerDelete(ctx context.Context, request *wwapi.ContainerDeleteParameter) (response *emptypb.Empty, err error) {

	response = new(emptypb.Empty)
	log.Println("ContainerDelete start")
	log.Printf("request: %T, %#v\n", request, request)

	// Parameter checks.
	if request == nil {
		return response, status.Errorf(codes.InvalidArgument, "nil request")
	}

	if request.ContainerNames == nil {
		return response, status.Errorf(codes.InvalidArgument, "nil request.ContainerNames")
	}

	err = container.ContainerDelete(request)
	return
}

func (s *apiServer) ContainerImport(ctx context.Context, request *wwapi.ContainerImportParameter) (response *wwapi.ContainerListResponse, err error) {
	// TODO: Remove traces on PR. (here and across the interface)
	log.Println("ContainerImport start")
	log.Printf("request: %T, %#v\n", request, request)


	// TODO: Missing the import call here ...
	var containerName string
	containerName, err = container.ContainerImport(request)
	if err != nil {
		return
	}

	// TODO: Need a container list for a single container here.
	var containers []*wwapi.ContainerInfo
	containers, err = container.ContainerList()
	if err != nil {
		return
	}

	// Container name may have been shimmed in during import,
	// which is why ContainerList returns it.
	for i := 0; i < len(containers); i ++ {
		if containerName == containers[i].Name {
			response = &wwapi.ContainerListResponse{
				Containers: []*wwapi.ContainerInfo{containers[i]},
			}
			return
		}
	}
	return
}

// ContainerList returns details about containers.
func (s *apiServer) ContainerList(ctx context.Context, request *emptypb.Empty) (response *wwapi.ContainerListResponse, err error) {
	// TODO: Remove traces on PR. (here and across the interface)
	log.Println("ContainerList start")
	log.Printf("request: %T, %#v\n", request, request)

	var containers []*wwapi.ContainerInfo
	containers, err = container.ContainerList()
	if err != nil {
		return
	}

	response = &wwapi.ContainerListResponse{
		Containers: containers,
	}
	return
}

// NodeAdd adds one or more nodes for management by Warewulf and returns the added nodes.
func (s *apiServer) NodeAdd(ctx context.Context, request *wwapi.NodeAddParameter) (response *wwapi.NodeListResponse, err error) {

	// TODO: Remove traces on PR. (here and across the interface)
	log.Println("NodeAdd start")
	log.Printf("request: %T, %#v\n", request, request)

	// Parameter checks.
	if request == nil {
		return response, status.Errorf(codes.InvalidArgument, "nil request")
	}

	if request.NodeNames == nil {
		return response, status.Errorf(codes.InvalidArgument, "nil request.NodeNames")
	}
	
	// Add the node(s).
	err = node.NodeAdd(request)
	if err != nil {
		return
	}

	// Return the added nodes as per REST.
	return s.nodeListInternal(request.NodeNames)
}

// NodeDelete deletes one or more nodes for removal of management by Warewulf.
func (s *apiServer) NodeDelete(ctx context.Context, request *wwapi.NodeDeleteParameter) (response *emptypb.Empty, err error) {

	response = new(emptypb.Empty)
	log.Println("NodeDelete start")
	log.Printf("request: %T, %#v\n", request, request)

	// Parameter checks.
	if request == nil {
		return response, status.Errorf(codes.InvalidArgument, "nil request")
	}

	if request.NodeNames == nil {
		return response, status.Errorf(codes.InvalidArgument, "nil request.NodeNames")
	}

	err = node.NodeDelete(request)
	return
}

// NodeList returns details about zero or more nodes.
func (s *apiServer) NodeList(ctx context.Context, request *wwapi.NodeNames) (response *wwapi.NodeListResponse, err error) {

	log.Println("NodeList start")
	log.Printf("request: %T, %#v\n", request, request)

	// Parameter checks. request.NodeNames can be nil.
	if request == nil {
		return response, status.Errorf(codes.InvalidArgument, "nil request")
	}

	// Perform the list.
	return s.nodeListInternal(request.NodeNames)
}

// NodeSet updates fields for zero or more nodes and returns the updated nodes.
func (s *apiServer) NodeSet(ctx context.Context, request *wwapi.NodeSetParameter) (response *wwapi.NodeListResponse, err error) {

	// Parameter checks.
	if request == nil {
		return response, status.Errorf(codes.InvalidArgument, "nil request")
	}

	if request.NodeNames == nil {
		return response, status.Errorf(codes.InvalidArgument, "nil request.NodeNames")
	}

	log.Println("NodeSet start")
	log.Printf("request: %T, %#v\n", request, request)

	// Perform the NodeSet.
	err = node.NodeSet(request)
	if err != nil {
		return
	}

	// Return the updated nodes as per REST.
	return s.nodeListInternal(request.NodeNames)
}

func (s *apiServer) NodeStatus(ctx context.Context, request *wwapi.NodeNames) (response *wwapi.NodeStatusResponse, err error) {

	log.Println("NodeStatus start")
	log.Printf("request: %T, %#v\n", request, request)

	// Parameter checks. request.NodeNames can be nil.
	if request == nil {
		return response, status.Errorf(codes.InvalidArgument, "nil request")
	}

	return node.NodeStatus(request.NodeNames)
}

// Version returns the versions of the wwapi and warewulf as well as the api
// prefix for http routes.
func (s *apiServer) Version(ctx context.Context, request *emptypb.Empty) (response *wwapi.VersionResponse, err error) {

	response = &wwapi.VersionResponse{
		ApiPrefix: apiPrefix,
		ApiVersion: apiVersion,
		WarewulfVersion: version.GetVersion(),
	}
	return
}

// Private helpers.

// nodeListInternal calls NodeList and returns NodeListResponse.
// This does not contain parameter checks.
func (s *apiServer) nodeListInternal(nodeNames []string) (response *wwapi.NodeListResponse, err error) {

	var nodes []*wwapi.NodeInfo
	nodes, err = node.NodeList(nodeNames)
	if err != nil {
		return
	}

	response = &wwapi.NodeListResponse{
		Nodes: nodes,
	}
	return
}