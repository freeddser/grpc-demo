package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc-customer/pem"
	customerPB "grpc-customer/proto/customer"
	userPB "grpc-customer/proto/user"
	"grpc-customer/server"
	"log"
	"net"
)

var (
	tls      = flag.Bool("tls", true, "Connection uses TLS if true, else plain TCP")
	certFile = flag.String("cert_file", "", "The TLS cert file")
	keyFile  = flag.String("key_file", "", "The TLS key file")
	port     = flag.Int("port", 50055, "The server port")
)

func main() {
	//echo "" | openssl s_client  -connect localhost:50051
	opts, listener := initServer()
	//with TLS
	s := grpc.NewServer(opts...)
	//register PB
	customerPB.RegisterCustomerServer(s, &server.CustomerServer{})
	userPB.RegisterUserServer(s, &server.UserServer{})
	s.Serve(listener)
}

func initServer() ([]grpc.ServerOption, net.Listener) {
	flag.Parse()
	var opts []grpc.ServerOption
	if *tls {
		if *certFile == "" {
			*certFile = pem.Path("qraved_com.crt")
		}
		if *keyFile == "" {
			*keyFile = pem.Path("qraved_com.key")
		}
		creds, err := credentials.NewServerTLSFromFile(*certFile, *keyFile)
		if err != nil {
			log.Fatalf("Failed to generate credentials %v", err)
		}
		log.Println("Init TLS done.")
		opts = []grpc.ServerOption{grpc.Creds(creds)}
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} else {
		log.Printf("Server listen on port :%d\n", *port)
	}

	return opts, lis
}
