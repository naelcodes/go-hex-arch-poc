package application

import (
	"github.com/naelcodes/ab-backend/internal/core/domain"
	"github.com/naelcodes/ab-backend/internal/infrastructure/persistence/postgres"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type Application struct {
	customerRepository domain.ICustomerRepository
	//invoiceRepository  domain.IInvoiceRepository
	//paymentRepository  domain.IPaymentRepository
}

func (application *Application) Init(globalContext *types.GlobalContext) {
	// TODO

	customerRepository := &postgres.CustomerRepository{Database: globalContext.Database, Context: globalContext.Context}

	application.customerRepository = customerRepository

}
