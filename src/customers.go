package main

import (
	"net/http"

	c "github.com/delineateio/mimas/common"
	"github.com/jinzhu/gorm"
	_ "github.com/swaggo/swag/example/celler/httputil"
)

// Customer represents a customer within this specific domain
type Customer struct {
	gorm.Model
	Forename string `json:"forename" binding:"required"`
	Surname  string `json:"surname" binding:"required"`
}

func addCustomer(request *c.Request, response *c.Response) {
	customer := Customer{}
	err := request.Map(&customer)
	if err != nil {
		response.Code = http.StatusBadRequest
		return
	}
	err = NewCustomerRepository().CreateCustomer(&customer)
	if err != nil {
		response.Code = http.StatusServiceUnavailable
		return
	}
	response.Code = http.StatusCreated
}
