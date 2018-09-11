package client

import (
	"fmt"
	"golang.org/x/net/context"
	userPB "grpcD/proto/user"
	"io"
	"log"
)

func CreateUser(client userPB.UserClient, customer *userPB.UserRequest) {
	resp, err := client.CreateUser(context.Background(), customer)
	if err != nil {
		log.Fatalf("Could not create Customer: %v", err)
	}
	if resp.Success {
		log.Printf("A new Customer has been added with id: %d", resp.Id)
	}
}

func GetUsers(client userPB.UserClient, filter *userPB.UserFilter) string {
	// calling the streaming API
	stream, err := client.GetUsers(context.Background(), filter)
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
	return "xx"
}

func UserDemo() {
	// creates a new CustomerClient
	userClient := userPB.NewUserClient(Conn)
	fmt.Println("add user")
	user := &userPB.UserRequest{
		Id:   10001,
		Name: "gavin",
	}

	// Create a new customer
	CreateUser(userClient, user)

	user = &userPB.UserRequest{
		Id:   10002,
		Name: "jack",
	}

	// Create a new customer

	CreateUser(userClient, user)
	//Filter with an empty Keyword
	fmt.Println("get user by Filter")
	userFilter := &userPB.UserFilter{Id: "10002"}
	userFilter2 := &userPB.UserFilter{Keyword: "gavin"}
	GetUsers(userClient, userFilter)
	abc := GetUsers(userClient, userFilter2)
	fmt.Println(abc)
}
