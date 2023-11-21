package restAdapter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naelcodes/ab-backend/internal/modules/imputations/ports"
	"github.com/naelcodes/ab-backend/internal/pkg/server"
)

type ImputationRestController struct {
	Application ports.IImputationApplication
}

func (controller *ImputationRestController) Init(appEngine *server.AppEngine) {

	//TODO : Replace this with global config
	baseRouter := appEngine.Get().Group("/api/v1")

	//validation middlewares here
	controller.attachImputationRoutesHandlers(baseRouter.Group("/invoice"))

}

func (controller *ImputationRestController) attachImputationRoutesHandlers(router fiber.Router) {
	router.Get("/id/imputations", controller.GetInvoiceImputationsHandler())
	router.Post("/id/imputations", controller.ApplyImputationsHandler())
	router.Patch("/id/imputations", controller.UpdateInvoiceImputationsHandler())
}
