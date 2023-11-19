package restAdapter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naelcodes/ab-backend/internal/modules/customers/ports"
	"github.com/naelcodes/ab-backend/internal/pkg/server"
)

type CostumerRestController struct {
	Application ports.ICustomerApplication
}

func (controller *CostumerRestController) Init(appEngine *server.AppEngine) {

	//TODO : Replace this with global config
	baseRouter := appEngine.Get().Group("/api/v1")

	//validation middlewares here
	controller.attachCustomerRoutesHandlers(baseRouter.Group("/customers"))
	controller.attachCountriesRouteHandler(baseRouter.Group("/countries"))
}

func (controller *CostumerRestController) attachCustomerRoutesHandlers(router fiber.Router) {

	router.Get("", controller.GetAllCustomersHandler())
	router.Get("/:id", controller.GetCustomerHandler())
	router.Post("", controller.CreateCustomerHandler())
	router.Put("/:id", controller.UpdateCustomerHandler())
	router.Delete("/:id", controller.deleteCustomerHandler())
}

func (controller *CostumerRestController) attachCountriesRouteHandler(router fiber.Router) {
	router.Get("", controller.GetAllCountriesHandler())
}
