package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-gateway-demo/gen/go/user"
	"log"
	"net"
)
// GreeterServerImpl will implement the service defined in protocol buffer definitions
type YourServiceImpl struct {
	grpc_gateway_demo.UnimplementedYourServiceServer
}
// SayHello is the implementation of RPC call defined in protocol definitions.
// This will take HelloRequest message and return HelloReply
func (g *YourServiceImpl) Echo(ctx context.Context, request *grpc_gateway_demo.StringMessage) (*grpc_gateway_demo.StringMessage, error) {

	return &grpc_gateway_demo.StringMessage{
		Value: "test",
	},nil
}
func main() {
	// create new gRPC server
	server := grpc.NewServer()
	// register the GreeterServerImpl on the gRPC server
	grpc_gateway_demo.RegisterYourServiceServer(server, &YourServiceImpl{})
	// start listening on port :8080 for a tcp connection
	if l, err := net.Listen("tcp", ":8081"); err != nil {
		log.Fatal("error in listening on port :8081", err)
	} else {
		// the gRPC server
		if err:=server.Serve(l);err!=nil {
			log.Fatal("unable to start server",err)
		}
	}

}