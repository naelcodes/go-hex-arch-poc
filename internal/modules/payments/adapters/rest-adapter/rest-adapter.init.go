package restAdapter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naelcodes/ab-backend/internal/modules/payments/ports"
	"github.com/naelcodes/ab-backend/internal/pkg/server"
)

type PaymentRestController struct {
	Application ports.IPaymentApplication
}

func (controller *PaymentRestController) Init(appEngine *server.AppEngine) {

	//TODO : Replace this with global config
	baseRouter := appEngine.Get().Group("/api/v1")

	//validation middlewares here
	controller.attachPaymentRoutesHandlers(baseRouter.Group("/payments"))
	controller.attachCustomerPaymentRouteHandler(baseRouter.Group("/customers"))
}

func (controller *PaymentRestController) attachPaymentRoutesHandlers(router fiber.Router) {
	router.Get("", controller.GetAllPaymentHandler())
	router.Post("", controller.AddCustomerPaymentHandler())
	router.Patch("/id", controller.UpdatePaymentHandler())
}

func (controller *PaymentRestController) attachCustomerPaymentRouteHandler(router fiber.Router) {
	router.Get("/id/payments", controller.GetCustomerPaymentHandler())
}
