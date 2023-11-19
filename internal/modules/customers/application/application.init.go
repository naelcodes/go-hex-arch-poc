package application

import (
	"github.com/naelcodes/ab-backend/internal/modules/customers/ports"
)

type CustomerApplication struct {
	repository ports.ICustomerRepository
}

func (application *CustomerApplication) Init(customerRepository ports.ICustomerRepository) {
	application.repository = customerRepository
}
