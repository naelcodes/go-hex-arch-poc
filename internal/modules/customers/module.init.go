package customers

import (
	"github.com/naelcodes/ab-backend/internal/common"
	RestAdpater "github.com/naelcodes/ab-backend/internal/modules/customers/adapters/rest-adapter"
	"github.com/naelcodes/ab-backend/internal/modules/customers/application"
)

func Init(globalContext *common.GlobalContext) {

	customerApplication := new(application.CustomerApplication)
	customerApplication.Init(globalContext)

	customerRestController := new(RestAdpater.CostumerRestController)
	customerRestController.Application = customerApplication

	customerRestController.Init(globalContext.AppEngine)
}
