package server

import (
	"golang.org/x/net/context"
	customerPB "grpcD/proto/customer"

	"strings"
)

// CustomerServer is used to implement customer.CustomerServer.
type CustomerServer struct {
	savedCustomers []*customerPB.CustomerRequest
}

// CreateCustomer creates a new Customer
func (s *CustomerServer) CreateCustomer(ctx context.Context, in *customerPB.CustomerRequest) (*customerPB.CustomerResponse, error) {
	s.savedCustomers = append(s.savedCustomers, in)
	return &customerPB.CustomerResponse{Id: in.Id, Success: true}, nil
}

// GetCustomers returns all customers by given filter
func (s *CustomerServer) GetCustomers(filter *customerPB.CustomerFilter, stream customerPB.Customer_GetCustomersServer) error {
	for _, customer := range s.savedCustomers {
		if filter.Keyword != "" {
			if !strings.Contains(customer.Name, filter.Keyword) {
				continue
			}
		}
		if err := stream.Send(customer); err != nil {
			return err
		}
	}
	return nil
}
