package server

import (
	"fmt"
	"golang.org/x/net/context"
	userPB "grpcD/proto/user"
	"grpcD/util"
	"strings"
)

type UserServer struct {
	savedUsers []*userPB.UserRequest
}

// CreateCustomer creates a new Customer
func (s *UserServer) CreateUser(ctx context.Context, in *userPB.UserRequest) (*userPB.UserResponse, error) {
	s.savedUsers = append(s.savedUsers, in)
	fmt.Println("save in user table")
	return &userPB.UserResponse{Id: in.Id, Success: true}, nil
}

// GetCustomers returns all customers by given filter
func (s *UserServer) GetUsers(filter *userPB.UserFilter, stream userPB.User_GetUsersServer) error {
	fmt.Println("get data from db")
	for _, user := range s.savedUsers {

		if filter.Keyword != "" {
			if !strings.Contains(user.Name, filter.Keyword) {
				continue
			}
		}

		if filter.Id != "" {
			if util.ConvertInt32ToString(user.Id) != filter.Id {
				continue
			}
		}

		if err := stream.Send(user); err != nil {
			return err
		}
	}
	return nil
}
