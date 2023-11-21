package restAdapter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naelcodes/ab-backend/internal/modules/invoices/ports"
	"github.com/naelcodes/ab-backend/internal/pkg/server"
)

type InvoiceRestController struct {
	Application ports.IInvoiceApplication
}

func (controller *InvoiceRestController) Init(appEngine *server.AppEngine) {

	//TODO : Replace this with global config
	baseRouter := appEngine.Get().Group("/api/v1")

	//validation middlewares here

	controller.attachInvoiceHandlers(baseRouter.Group("/invoices"))
	controller.attachTravelItemsHandler(baseRouter.Group("/travel_items"))
}

func (controller *InvoiceRestController) attachInvoiceHandlers(router fiber.Router) {

	router.Get("", controller.GetAllInvoiceHandler())
	router.Get("/:id", controller.GetInvoiceHandler())
	router.Post("", controller.CreateInvoiceHandler())
	router.Put("/:id", controller.UpdateInvoiceHandler())
	router.Delete("/:id", controller.DeleteInvoiceHandler())
}

func (controller *InvoiceRestController) attachTravelItemsHandler(router fiber.Router) {
	router.Get("", controller.GetAllTravelItemsHandler())
}
