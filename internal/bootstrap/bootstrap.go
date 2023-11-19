package bootstrap

import (
	"github.com/naelcodes/ab-backend/internal/common"
	customersModule "github.com/naelcodes/ab-backend/internal/modules/customers"
)

// Inject module dependencies
func InitModules(globalContext *common.GlobalContext) {

	customersModule.Init(globalContext)

	// invoiceModule := new(invoiceRestAdapter.InvoiceModule)
	// invoiceRepository := invoiceDatabaseAdapter.InvoiceRepository{}
	// invoiceModule.Init(appEngine, invoiceRepository)

}
