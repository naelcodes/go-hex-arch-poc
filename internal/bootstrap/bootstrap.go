package bootstrap

import (
	"github.com/naelcodes/ab-backend/internal/common"
	customerModule "github.com/naelcodes/ab-backend/internal/modules/customers"
	// imputationModule "github.com/naelcodes/ab-backend/internal/modules/imputations"
	// invoiceModule "github.com/naelcodes/ab-backend/internal/modules/invoices"
	// paymentModule "github.com/naelcodes/ab-backend/internal/modules/payments"
)

// Inject module dependencies
func InitModules(globalContext *common.GlobalContext) {

	customerModule.Init(globalContext)
	// invoiceModule.Init(globalContext)
	// paymentModule.Init(globalContext)
	// imputationModule.Init(globalContext)

}
