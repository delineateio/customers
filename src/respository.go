package main

import (
	c "github.com/delineateio/mimas/common"
)

// CustomerRepository that reprents the access to the underlying database
type CustomerRepository struct {
	core *c.Repository
}

// NewCustomerRepository returns production database access
func NewCustomerRepository() *CustomerRepository {
	return &CustomerRepository{
		core: c.NewRepository("customers"),
	}
}

// Migrate the DB to the latest schema
func (customers *CustomerRepository) Migrate() error {
	// Migrates the model
	return customers.core.Migrate(&Customer{})
}

// CreateCustomer adds the customer object to the database
func (customers *CustomerRepository) CreateCustomer(customer *Customer) error {
	// Creates the customer
	return customers.core.Create(&customer)
}
