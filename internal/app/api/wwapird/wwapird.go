package main

import (
	"context"
	//"flag"
	"fmt"
	//"os"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/hpcng/warewulf/internal/pkg/api/apiconfig"

	gw "github.com/hpcng/warewulf/internal/pkg/api/routes/wwapiv1"

	"github.com/hpcng/warewulf/internal/pkg/buildconfig"
	"path"
)

//var (
	// command-line options:
	// gRPC server endpoint
	// TODO: Config file (port, etc)
	//grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9872", "gRPC server endpoint")
//)

func run() error {

	// Read the config file.
	config, err := apiconfig.NewClientServer(path.Join(buildconfig.SYSCONFDIR(), "warewulf/wwapird.conf"))
	if err != nil {
		glog.Fatalf("Failed to read config file, err: %v", err)
	}

	grpcServerEndpoint := fmt.Sprintf("%s:%v", config.ClientApiConfig.Server, config.ClientApiConfig.Port)
	
	

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint (we are the client)
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())} // TODO: This is the sever side change.
	//err := gw.RegisterYourServiceHandlerFromEndpoint(ctx, mux,  *grpcServerEndpoint, opts)
	//err := gw.RegisterWWApiHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts)
	err = gw.RegisterWWApiHandlerFromEndpoint(ctx, mux, grpcServerEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":9871", mux) // TODO: This is the client side change.
}

func main() {
	//flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}