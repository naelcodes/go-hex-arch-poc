package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/naelcodes/ab-backend/internal/core/application"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	"github.com/naelcodes/ab-backend/internal/infrastructure/api/middleware"
	"github.com/naelcodes/ab-backend/pkg/types"
	"github.com/naelcodes/ab-backend/pkg/utils"
)

type RestController struct {
	ApplicationService *application.Application
}

func (controller *RestController) Init(globalContext *types.GlobalContext) {
	server := globalContext.AppEngine.GetServer()

	server.Use(recover.New())
	server.Use(middleware.Cors())
	server.Use(middleware.QueryValidator())

	api := server.Group("/api")
	v1 := api.Group("/v1")

	customerSubApp := controller.CustomerSubApp()
	travelItemSubApp := controller.TravelItemSubApp()
	invoiceSubApp := controller.InvoiceSubApp()
	paymentSubApp := controller.PaymentSubApp()
	invoiceImputationSubApp := controller.InvoiceImputationSubApp()

	utils.Logger.Info("[Init] - Mounting sub-apps")

	v1.Mount("/customers", customerSubApp)
	utils.Logger.Info(fmt.Sprintf("[Init] - customer sub-app mount path: %v", customerSubApp.MountPath()))

	v1.Mount("/travel-items", travelItemSubApp)
	utils.Logger.Info(fmt.Sprintf("[Init] - travel item sub-app mount path: %v", travelItemSubApp.MountPath()))

	v1.Mount("/invoices", invoiceSubApp)
	utils.Logger.Info(fmt.Sprintf("[Init] - invoice sub-app mount path: %v", invoiceSubApp.MountPath()))

	v1.Mount("/invoices", invoiceImputationSubApp)
	utils.Logger.Info(fmt.Sprintf("[Init] - invoice imputation sub-app mount path: %v", invoiceImputationSubApp.MountPath()))

	v1.Mount("/payments", paymentSubApp)
	utils.Logger.Info(fmt.Sprintf("[Init] - payment sub-app mount path: %v", paymentSubApp.MountPath()))

}

func (controller *RestController) CustomerSubApp() *fiber.App {

	micro := fiber.New()

	micro.Use(middleware.PayloadValidator(new(dto.CreateCustomerDTO), new(dto.UpdateCustomerDTO)))
	micro.Get("/", controller.GetAllCustomersHandler())
	micro.Get("/:id", controller.GetCustomerHandler())
	micro.Get("/:id/invoices", controller.GetCustomerInvoicesHandler())
	micro.Post("", controller.CreateCustomerHandler())
	micro.Patch("/:id", controller.UpdateCustomerHandler())
	micro.Delete("/:id", controller.DeleteCustomerHandler())
	return micro
}

func (controller *RestController) TravelItemSubApp() *fiber.App {
	micro := fiber.New()
	micro.Get("", controller.GetAllTravelItemsHandler())
	return micro
}

func (controller *RestController) InvoiceSubApp() *fiber.App {

	micro := fiber.New()
	micro.Use(middleware.PayloadValidator(new(dto.CreateInvoiceDTO), new(dto.UpdateInvoiceDTO)))
	micro.Get("", controller.GetAllInvoiceHandler())
	micro.Get("/:id", controller.GetInvoiceHandler())
	micro.Post("", controller.CreateInvoiceHandler())
	micro.Patch("/:id", controller.UpdateInvoiceHandler())
	micro.Delete("/:id", controller.DeleteInvoiceHandler())

	return micro
}

func (controller *RestController) InvoiceImputationSubApp() *fiber.App {

	micro := fiber.New()
	micro.Use(middleware.ImputationPayloadValidator(make([]*dto.InvoiceImputationDTO, 0)))
	micro.Get("/:id/imputations", controller.GetInvoiceImputationsHandler())
	micro.Post("/:id/imputations", controller.ApplyInvoiceImputationHandler())

	return micro

}

func (controller *RestController) PaymentSubApp() *fiber.App {

	micro := fiber.New()

	micro.Use(middleware.PayloadValidator(new(dto.CreatePaymentDTO), new(dto.UpdatePaymentDTO)))
	micro.Get("", controller.GetAllPaymentsHandler())
	micro.Get("/:id", controller.GetPaymentHandler())
	micro.Post("", controller.CreatePaymentHandler())
	micro.Patch("/:id", controller.UpdatePaymentHandler())
	micro.Delete("/:id", controller.DeletePaymentHandler())

	return micro
}
