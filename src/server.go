package main

import (
	messages "github.com/delineateio/mimas/messages"
	server "github.com/delineateio/mimas/server"
)

func main() {
	// Defines the routes for the service
	routes := []messages.Route{
		{Method: "POST", Path: "/customer", Handler: addCustomer},
	}

	instance := server.NewServer(routes)
	instance.Repository = NewCustomerRepository()
	instance.Start()
}
