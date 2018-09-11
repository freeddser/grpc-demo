package server

import (
	"fmt"
	"golang.org/x/net/context"
	userPB "grpc-customer/proto/user"
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
			fmt.Println("xxx")
			if !strings.Contains(user.Name, filter.Keyword) {
				continue
			}
		}

		if filter.Id != "" {
			if String(user.Id) != filter.Id {
				continue
			}
		}

		if err := stream.Send(user); err != nil {
			return err
		}
	}
	return nil
}

func String(n int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(n)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}
