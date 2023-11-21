package application

import (
	"context"

	"github.com/naelcodes/ab-backend/internal/modules/customers/ports"
)

type CustomerApplication struct {
	context    context.Context
	repository ports.ICustomerRepository
}

func (application *CustomerApplication) Init(context context.Context, customerRepository ports.ICustomerRepository) {
	application.repository = customerRepository
	application.context = context
}
