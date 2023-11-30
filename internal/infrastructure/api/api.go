package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/naelcodes/ab-backend/internal/core/application"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	"github.com/naelcodes/ab-backend/internal/infrastructure/api/middleware"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type RestController struct {
	ApplicationService *application.Application
}

func (controller *RestController) Init(globalContext *types.GlobalContext) {
	server := globalContext.AppEngine.GetServer()

	server.Use(recover.New())
	server.Use(middleware.Cors())
	server.Use(middleware.QueryValidator())

	apiV1Router := server.Group("/api/v1")
	controller.attachCustomerRoutesHandlers(apiV1Router.Group("/customers"))
	controller.attachTravelItemsRoutesHandlers(apiV1Router.Group("/travel-items"))
	controller.attachInvoiceRoutesHandlers(apiV1Router.Group("/invoices"))
	controller.attachInvoiceImputationRoutesHandlers(apiV1Router.Group("/invoices"))
	controller.attachPaymentRoutesHandlers(apiV1Router.Group("/payments"))
}

func (controller *RestController) attachCustomerRoutesHandlers(router fiber.Router) {

	router.Use(middleware.PayloadValidator(new(dto.CreateCustomerDTO), new(dto.UpdateCustomerDTO)))
	router.Get("", controller.GetAllCustomersHandler())
	router.Get("/:id", controller.GetCustomerHandler())
	// router.Get("/:id/payments", controller.GetCustomerPaymentsHandler())
	// router.Get("/:id/invoices", controller.GetCustomerInvoicesHandler())
	router.Post("", controller.CreateCustomerHandler())
	router.Patch("/:id", controller.UpdateCustomerHandler())
	router.Delete("/:id", controller.DeleteCustomerHandler())
}

func (controller *RestController) attachTravelItemsRoutesHandlers(router fiber.Router) {
	router.Get("", controller.GetAllTravelItemsHandler())
}

func (controller *RestController) attachInvoiceRoutesHandlers(router fiber.Router) {
	router.Get("", controller.GetAllInvoiceHandler())
	router.Get("/:id", controller.GetInvoiceHandler())
	router.Post("", controller.CreateInvoiceHandler())
	router.Patch("/:id", controller.UpdateInvoiceHandler())
	router.Delete("/:id", controller.DeleteInvoiceHandler())
}

func (controller *RestController) attachInvoiceImputationRoutesHandlers(router fiber.Router) {

	router.Get("/:id/imputations", controller.GetInvoiceImputationsHandler())
	router.Post(":id/imputations", controller.AddInvoiceImputationHandler())
	router.Patch("/:id/imputations", controller.UpdateInvoiceImputationsHandler())
}

func (controller *RestController) attachPaymentRoutesHandlers(router fiber.Router) {

	router.Get("", controller.GetAllPaymentsHandler())
	router.Get("/:id", controller.GetPaymentHandler())
	router.Post("", controller.CreatePaymentHandler())
	router.Patch("/:id", controller.UpdatePaymentHandler())
	router.Delete("/:id", controller.DeletePaymentHandler())

}
