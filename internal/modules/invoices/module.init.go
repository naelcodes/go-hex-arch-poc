package invoices

import (
	"github.com/naelcodes/ab-backend/internal/common"
	PostgresAdapter "github.com/naelcodes/ab-backend/internal/modules/invoices/adapters/postgres-adapter"
	RestAdapter "github.com/naelcodes/ab-backend/internal/modules/invoices/adapters/rest-adapter"
	"github.com/naelcodes/ab-backend/internal/modules/invoices/application"
)

func Init(globalContext *common.GlobalContext) {
	invoiceRepository := &PostgresAdapter.InvoiceRepository{Database: globalContext.Database}

	invoiceApplication := new(application.InvoiceApplication)
	invoiceApplication.Init(invoiceRepository)

	invoiceRestController := new(RestAdapter.InvoiceRestController)
	invoiceRestController.Application = invoiceApplication

	invoiceRestController.Init(globalContext.AppEngine)

}
