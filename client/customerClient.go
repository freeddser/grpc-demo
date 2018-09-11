package client

import (
	"golang.org/x/net/context"
	customerPB "grpcD/proto/customer"
	"io"
	"log"
)

// createCustomer calls the RPC method CreateCustomer of CustomerServer
func CreateCustomer(client customerPB.CustomerClient, customer *customerPB.CustomerRequest) {
	resp, err := client.CreateCustomer(context.Background(), customer)
	if err != nil {
		log.Fatalf("Could not create Customer: %v", err)
	}
	if resp.Success {
		log.Printf("A new Customer has been added with id: %d", resp.Id)
	}
}

// GetCustomers calls the RPC method GetCustomers of CustomerServer
func GetCustomers(client customerPB.CustomerClient, filter *customerPB.CustomerFilter) {
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

func CustomerDemo() {
	// creates a new CustomerClient
	customerclient := customerPB.NewCustomerClient(Conn)
	//
	customer := &customerPB.CustomerRequest{
		Id:    101,
		Name:  "Shiju Varghese",
		Email: "shiju@xyz.com",
		Phone: "732-757-2923",
		Addresses: []*customerPB.CustomerRequest_Address{
			&customerPB.CustomerRequest_Address{
				Street:            "1 Mission Street",
				City:              "San Francisco",
				State:             "CA",
				Zip:               "94105",
				IsShippingAddress: false,
			},
			&customerPB.CustomerRequest_Address{
				Street:            "Greenfield",
				City:              "Kochi",
				State:             "KL",
				Zip:               "68356",
				IsShippingAddress: true,
			},
		},
	}

	//// Create a new customer
	CreateCustomer(customerclient, customer)

	customer = &customerPB.CustomerRequest{
		Id:    102,
		Name:  "Irene Rose",
		Email: "irene@xyz.com",
		Phone: "732-757-2924",
		Addresses: []*customerPB.CustomerRequest_Address{
			&customerPB.CustomerRequest_Address{
				Street:            "1 Mission Street",
				City:              "San Francisco",
				State:             "CA",
				Zip:               "94105",
				IsShippingAddress: true,
			},
		},
	}

	//// Create a new customer
	CreateCustomer(customerclient, customer)
	////Filter with an empty Keyword
	filter := &customerPB.CustomerFilter{Keyword: "Irene Rose"}
	GetCustomers(customerclient, filter)
}
