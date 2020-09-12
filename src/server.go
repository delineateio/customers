package main

import (
	"github.com/delineateio/mimas/handlers"
	"github.com/delineateio/mimas/msgs"
	"github.com/delineateio/mimas/server"
	"gorm.io/gorm"
)

// Customer represents a customer within this specific domain
type Customer struct {
	gorm.Model
	Forename string `json:"forename" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
}

func main() {
	opts, _ := server.NewDefaultOptions()
	opts.AddRoute("POST", "/customer", createCustomer)
	opts.AddEntities(&Customer{})
	s := server.NewServer(opts)
	s.Listen()
}

func createCustomer(request *msgs.Request, response *msgs.Response) {
	var entity interface{} = &Customer{}
	handlers.Create(request, entity, response)
}
