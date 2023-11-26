package application

import "github.com/naelcodes/ab-backend/internal/core/domain"

type Application struct {
	customerRepository domain.ICustomerRepository
	//invoiceRepository  domain.IInvoiceRepository
	//paymentRepository  domain.IPaymentRepository
}

func (application *Application) Init() {
	// TODO
}
