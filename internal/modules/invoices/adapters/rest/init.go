package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naelcodes/ab-backend/internal/modules/invoices/domain"
	"github.com/naelcodes/ab-backend/internal/pkg/server"
)

type InvoiceModule struct {
	Repository domain.IInvoiceRepository
}

func (im *InvoiceModule) Init(appEngine *server.AppEngine, repository domain.IInvoiceRepository) {

	//TODO : Replace this with global config
	apiGroup := appEngine.Get().Group("/api/v1")

	//validation middlewares here

	im.Repository = repository

	im.attachInvoiceHandlers(apiGroup.Group("/invoices"))
	im.attachTravelItemsHandler(apiGroup.Group("/travel_items"))
}

func (im *InvoiceModule) attachInvoiceHandlers(router fiber.Router) {

	router.Get("", im.GetAllInvoiceHandler())
	router.Get("/:id", im.GetInvoiceHandler())
	router.Post("", im.CreateInvoiceHandler())
	router.Put("/:id", im.UpdateInvoiceHandler())
	router.Delete("/:id", im.deleteInvoiceHandler())
}

func (im *InvoiceModule) attachTravelItemsHandler(router fiber.Router) {
	router.Get("", im.GetAllTravelItemsHandler())
}
