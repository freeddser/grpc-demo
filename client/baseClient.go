package client

import (
	"log"

	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"grpc-customer/pem"
	"os"
)

var (
	tls                = flag.Bool("tls", true, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containning the CA root cert file")
	serverAddr         = flag.String("server_addr", "127.0.0.1:55555", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "www.lijiuyang.com", "The server name use to verify the hostname returned by TLS handshake")
)

var Conn *grpc.ClientConn

func InitConn() {
	var err error
	if Conn == nil {
		Conn, err = GetClientConn()
		if err != nil {
			log.Fatalf("Failed to get Conn %v", err)
			os.Exit(1)
		}
	}
}

func GetClientConn() (*grpc.ClientConn, error) {
	flag.Parse()
	var opts []grpc.DialOption
	log.Println("init ssl and connect to RPC server...")
	if *tls {
		if *caFile == "" {
			*caFile = pem.Path("www.lijiuyang.com.crt")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)

		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
			return nil, err
		}
		log.Printf("Init Client CA:%s done.", *caFile)
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatal("did not connect: %v", err)
		return nil, err
	}
	return conn, nil
}
