// client/main.go
package main

import (
	"google.golang.org/grpc"
	"grpc-customer/client"
)

var Conn *grpc.ClientConn

func main() {
	//get client rpc connect
	client.InitConn()

	//run client demo
	client.UserDemo()
	client.CustomerDemo()
}
