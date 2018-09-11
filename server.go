package main

import (
	"google.golang.org/grpc"
	customerPB "grpcD/proto/customer"
	userPB "grpcD/proto/user"
	"grpcD/server"
)

func main() {
	//echo "" | openssl s_client  -connect localhost:50051
	opts, listener := server.InitServer()
	//with TLS
	s := grpc.NewServer(opts...)

	//register PB
	customerPB.RegisterCustomerServer(s, &server.CustomerServer{})
	userPB.RegisterUserServer(s, &server.UserServer{})
	s.Serve(listener)
}
