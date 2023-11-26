package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naelcodes/ab-backend/internal/core/application"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type RestController struct {
	ApplicationService *application.Application
}

func (controller *RestController) Init(globalContext *types.GlobalContext) {
	appEngine := globalContext.AppEngine

	baseRouter := appEngine.GetServer().Group("/api/v1")
	//validation middlewares here
	controller.attachCustomerRoutesHandlers(baseRouter.Group("/customers"))
}

func (controller *RestController) attachCustomerRoutesHandlers(router fiber.Router) {

	router.Get("", controller.GetAllCustomersHandler())
	router.Get("/:id", controller.GetCustomerHandler())
	router.Post("", controller.CreateCustomerHandler())
	router.Put("/:id", controller.UpdateCustomerHandler())
	router.Delete("/:id", controller.DeleteCustomerHandler())
}
