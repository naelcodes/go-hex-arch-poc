package customers

import (
	"github.com/naelcodes/ab-backend/internal/common"
	PostgresAdapter "github.com/naelcodes/ab-backend/internal/modules/customers/adapters/postgres-adapter"
	RestAdpater "github.com/naelcodes/ab-backend/internal/modules/customers/adapters/rest-adapter"
	"github.com/naelcodes/ab-backend/internal/modules/customers/application"
)

func Init(globalContext *common.GlobalContext) {

	customerRepository := &PostgresAdapter.CustomerRepository{Database: globalContext.Database}

	customerApplication := new(application.CustomerApplication)
	customerApplication.Init(customerRepository)

	customerRestController := new(RestAdpater.CostumerRestController)
	customerRestController.Application = customerApplication

	customerRestController.Init(globalContext.AppEngine)
}
