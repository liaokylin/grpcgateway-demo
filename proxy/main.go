package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	gen "grpc-gateway-demo/gen/go/user"
	"log"
	"net"
	"net/http"
)

func main() {
	// creating mux for gRPC gateway. This will multiplex or route request different gRPC service
	mux := runtime.NewServeMux()
	// setting up a dail up for gRPC service by specifying endpoint/target url
	err := gen.RegisterGreeterHandlerFromEndpoint(context.Background(), mux, "localhost:8080", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatal(err)
	}
	err = gen.RegisterYourServiceHandlerFromEndpoint(context.Background(), mux, "localhost:8081", []grpc.DialOption{grpc.WithInsecure()})
	if err != nil {
		log.Fatal(err)
	}
	// Creating a normal HTTP server
	server := http.Server{
		Handler: mux,
	}
	// creating a listener for server
	l, err := net.Listen("tcp", ":8090")
	if err != nil {
		log.Fatal(err)
	}
	// start server
	err = server.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}
