// client/main.go
package main

import (
	"io"
	"log"

	pb "customer"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"testdata"
    	"flag"
)


var (
	tls                = flag.Bool("tls", true, "Connection uses TLS if true, else plain TCP")
	caFile             = flag.String("ca_file", "", "The file containning the CA root cert file")
	serverAddr         = flag.String("server_addr", "127.0.0.1:55555", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "www.qraved.com", "The server name use to verify the hostname returned by TLS handshake")
)


// createCustomer calls the RPC method CreateCustomer of CustomerServer
func createCustomer(client pb.CustomerClient, customer *pb.CustomerRequest) {
	resp, err := client.CreateCustomer(context.Background(), customer)
	if err != nil {
		log.Fatalf("Could not create Customer: %v", err)
	}
	if resp.Success {
		log.Printf("A new Customer has been added with id: %d", resp.Id)
	}
}

// GetCustomers calls the RPC method GetCustomers of CustomerServer
func getCustomers(client pb.CustomerClient, filter *pb.CustomerFilter) {
	// calling the streaming API
	stream, err := client.GetCustomers(context.Background(), filter)
	if err != nil {
		log.Fatal("Error on get customers: %v", err)
	}
	for {
		customer, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal("%v.GetCustomers(_) = _, %v", client, err)
		}
		log.Printf("Customer: %v", customer)
	}
}

func main() {

	flag.Parse()
	var opts []grpc.DialOption

	if *tls {
		if *caFile == "" {
			*caFile = testdata.Path("qraved_com.crt")
		}
		creds, err := credentials.NewClientTLSFromFile(*caFile, *serverHostOverride)

		if err != nil {
			log.Fatalf("Failed to create TLS credentials %v", err)
		}
		log.Printf("Init Client CA:%s done.",*caFile)
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithInsecure())
	}
	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()
	// creates a new CustomerClient
	client := pb.NewCustomerClient(conn)

	customer := &pb.CustomerRequest{
		Id:    101,
		Name:  "Shiju Varghese",
		Email: "shiju@xyz.com",
		Phone: "732-757-2923",
		Addresses: []*pb.CustomerRequest_Address{
			&pb.CustomerRequest_Address{
				Street:            "1 Mission Street",
				City:              "San Francisco",
				State:             "CA",
				Zip:               "94105",
				IsShippingAddress: false,
			},
			&pb.CustomerRequest_Address{
				Street:            "Greenfield",
				City:              "Kochi",
				State:             "KL",
				Zip:               "68356",
				IsShippingAddress: true,
			},
		},
	}

	// Create a new customer
	createCustomer(client, customer)

	customer = &pb.CustomerRequest{
		Id:    102,
		Name:  "Irene Rose",
		Email: "irene@xyz.com",
		Phone: "732-757-2924",
		Addresses: []*pb.CustomerRequest_Address{
			&pb.CustomerRequest_Address{
				Street:            "1 Mission Street",
				City:              "San Francisco",
				State:             "CA",
				Zip:               "94105",
				IsShippingAddress: true,
			},
		},
	}

	// Create a new customer
	createCustomer(client, customer)
	//Filter with an empty Keyword
	filter := &pb.CustomerFilter{Keyword: "Irene Rose"}
	getCustomers(client, filter)

}
