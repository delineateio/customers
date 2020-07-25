package main

import (
	c "github.com/delineateio/mimas/common"
	s "github.com/delineateio/mimas/server"
)

func main() {
	// Defines the routes for the service
	routes := []c.Route{
		{Method: "POST", Path: "/customer", Handler: addCustomer},
	}

	server := s.NewServer(routes)
	server.Repository = NewCustomerRepository()
	server.Start()
}
