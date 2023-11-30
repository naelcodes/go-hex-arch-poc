package application

import (
	customerDomain "github.com/naelcodes/ab-backend/internal/core/domains/customer-domain"
	invoiceDomain "github.com/naelcodes/ab-backend/internal/core/domains/invoice-domain"
	paymentDomain "github.com/naelcodes/ab-backend/internal/core/domains/payment-domain"
	"github.com/naelcodes/ab-backend/internal/infrastructure/persistence/postgres"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type Application struct {
	customerRepository customerDomain.ICustomerRepository
	invoiceRepository  invoiceDomain.IInvoiceRepository
	paymentRepository  paymentDomain.IPaymentRepository
}

func (application *Application) Init(globalContext *types.GlobalContext) {
	// TODO

	customerRepository := &postgres.CustomerRepository{Database: globalContext.Database, Context: globalContext.Context}

	// invoiceRepository := &postgres.InvoiceRepository{Database: globalContext.Database, Context: globalContext.Context}

	// paymentRepository := &postgres.PaymentRepository{Database: globalContext.Database, Context: globalContext.Context}

	application.customerRepository = customerRepository
	// application.invoiceRepository = invoiceRepository
	// application.paymentRepository = paymentRepository

}
