package application

import (
	travelItemDomain "github.com/naelcodes/ab-backend/internal/core/domains/TravelItem-domain"
	customerDomain "github.com/naelcodes/ab-backend/internal/core/domains/customer-domain"
	imputationDomain "github.com/naelcodes/ab-backend/internal/core/domains/imputation-domain"
	invoiceDomain "github.com/naelcodes/ab-backend/internal/core/domains/invoice-domain"
	paymentDomain "github.com/naelcodes/ab-backend/internal/core/domains/payment-domain"
	"github.com/naelcodes/ab-backend/internal/infrastructure/persistence/postgres"

	"github.com/naelcodes/ab-backend/pkg/logger"
	"github.com/naelcodes/ab-backend/pkg/transactionManager"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type Application struct {
	Logger               *logger.Logger
	customerRepository   customerDomain.ICustomerRepository
	invoiceRepository    invoiceDomain.IInvoiceRepository
	paymentRepository    paymentDomain.IPaymentRepository
	imputationRepository imputationDomain.IImputationRepository
	travelItemRepository travelItemDomain.ITravelItemRepository

	TransactionManager *transactionManager.TransactionManager
}

func (application *Application) Init(globalContext *types.GlobalContext) {

	application.Logger = globalContext.AppEngine.GetLogger()

	customerRepository := &postgres.CustomerRepository{Database: globalContext.Database, Context: globalContext.Context, Logger: globalContext.AppEngine.GetLogger()}
	application.customerRepository = customerRepository

	invoiceRepository := &postgres.InvoiceRepository{Database: globalContext.Database, Context: globalContext.Context, Logger: globalContext.AppEngine.GetLogger()}
	application.invoiceRepository = invoiceRepository

	paymentRepository := &postgres.PaymentRepository{Database: globalContext.Database, Context: globalContext.Context, Logger: globalContext.AppEngine.GetLogger()}
	application.paymentRepository = paymentRepository

	imputationRepository := &postgres.ImputationRepository{Database: globalContext.Database, Context: globalContext.Context}
	application.imputationRepository = imputationRepository

	travelItemRepository := &postgres.TravelItemRepository{Database: globalContext.Database, Context: globalContext.Context, Logger: globalContext.AppEngine.GetLogger()}
	application.travelItemRepository = travelItemRepository

	application.TransactionManager = transactionManager.NewTransactionManager(globalContext.Context, globalContext.Database)

}
