package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"log"
	"time"
	"io/ioutil"

	wwapi "github.com/hpcng/warewulf/internal/pkg/api/routes/wwapiv1"

	"google.golang.org/grpc"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/grpc/credentials"

	//"google.golang.org/grpc/credentials/insecure" // TODO: Block client on startup if in insecure mode.
)

// wwapic is intended as a sample wwapi client.

func main() {
	log.Println("Client running")

	/*
	// TODO: Remove hardcoded port. Config file. Also TLS is hardcoded.
	//certFile := "/home/mhink/mtls/client-cert.pem"
	//keyFile := "/home/mhink/mtls/client-key.pem"
	var opts []grpc.DialOption
	//certFile := "/home/mhink/mtls/client-crt.pem"
	caFile := "/home/mhink/mtls/ca-crt.pem"
	creds, err := credentials.NewClientTLSFromFile(caFile, "")
	if err != nil {
		log.Fatalf("Failed to create TLS credentials %v", err)
	}
	opts = append(opts, grpc.WithTransportCredentials(creds))

	//conn, err := grpc.Dial(":9872", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	//conn, err := grpc.Dial("rocky:9872", opts...) // TODO: Servername in config.
	*/

	// Load the client cert and its key
	clientCert, err := tls.LoadX509KeyPair("/home/mhink/mtls/client.pem", "/home/mhink/mtls/client.key")
	if err != nil {
		log.Fatalf("Failed to load client cert and key. %s.", err)
	}

	// Load the CA cert.
	var cacert []byte
	cacert, err = ioutil.ReadFile("/home/mhink/mtls/cacert.pem")
	if err != nil {
		log.Fatalf("Failed to load cacert. err: %s\n", err)
	}
	
	// Put the CA cert into the cert pool.
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(cacert) {
		log.Fatalf("Failed to append CA cert to certificate pool. %s.", err)
	}

	// Create the TLS configuration
	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
		MinVersion:   tls.VersionTLS13,
		MaxVersion:   tls.VersionTLS13,
	}

	// Create TLS credentials from the TLS configuration
	creds := credentials.NewTLS(tlsConfig)

		
	//conn, err := grpc.Dial("localhost:9872", opts...) // TODO: Servername in config.
	conn, err := grpc.Dial("localhost:9872", grpc.WithTransportCredentials(creds)) // TODO: Servername in config.
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

