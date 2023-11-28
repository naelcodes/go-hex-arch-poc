package api

import (
	"github.com/gofiber/fiber/v2"
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

	server.Use(middleware.Cors())
	server.Use(middleware.QueryValidator())

	apiV1Router := server.Group("/api/v1")
	controller.attachCustomerRoutesHandlers(apiV1Router.Group("/customers"))
	controller.attachTravelItemsRoutesHandlers(apiV1Router.Group("/travel-items"))
	controller.attachInvoiceRoutesHandlers(apiV1Router.Group("/invoices"))
	controller.attachPaymentRoutesHandlers(apiV1Router.Group("/payments"))
}

func (controller *RestController) attachCustomerRoutesHandlers(router fiber.Router) {

	router.Use(middleware.PayloadValidator(new(dto.CreateCustomerDTO), new(dto.UpdateCustomerDTO)))
	router.Get("", controller.GetAllCustomersHandler())
	router.Get("/:id", controller.GetCustomerHandler())
	router.Post("", controller.CreateCustomerHandler())
	router.Put("/:id", controller.UpdateCustomerHandler())
	router.Delete("/:id", controller.DeleteCustomerHandler())
}

func (controller *RestController) attachTravelItemsRoutesHandlers(router fiber.Router) {

}

func (controller *RestController) attachInvoiceRoutesHandlers(router fiber.Router) {}

func (controller *RestController) attachPaymentRoutesHandlers(router fiber.Router) {}
