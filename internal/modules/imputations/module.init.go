package imputations

import (
	"github.com/naelcodes/ab-backend/internal/common"
	PostgresAdapter "github.com/naelcodes/ab-backend/internal/modules/imputations/adapters/postgres-adapter"
	RestAdapter "github.com/naelcodes/ab-backend/internal/modules/imputations/adapters/rest-adapter"
	"github.com/naelcodes/ab-backend/internal/modules/imputations/application"
)

func Init(globalContext *common.GlobalContext) {

	imputationRepository := &PostgresAdapter.ImputationRepository{Database: globalContext.Database}

	imputationApplication := new(application.ImputationApplication)
	imputationApplication.Init(imputationRepository)

	ImputationRestController := new(RestAdapter.ImputationRestController)
	ImputationRestController.Application = imputationApplication

	ImputationRestController.Init(globalContext.AppEngine)
}
