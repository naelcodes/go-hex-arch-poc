package payments

import (
	"github.com/naelcodes/ab-backend/internal/common"
	PostgresAdapter "github.com/naelcodes/ab-backend/internal/modules/payments/adapters/postgres-adapter"
	RestAdapter "github.com/naelcodes/ab-backend/internal/modules/payments/adapters/rest-adapter"
	"github.com/naelcodes/ab-backend/internal/modules/payments/application"
)

func Init(globalContext *common.GlobalContext) {

	paymentRepository := &PostgresAdapter.PaymentRepository{Database: globalContext.Database}

	paymentApplication := new(application.PaymentApplication)
	paymentApplication.Init(paymentRepository)

	paymentRestController := new(RestAdapter.PaymentRestController)
	paymentRestController.Application = paymentApplication

	paymentRestController.Init(globalContext.AppEngine)

}
