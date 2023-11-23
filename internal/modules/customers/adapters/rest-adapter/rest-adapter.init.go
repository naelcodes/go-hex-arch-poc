package restAdapter

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naelcodes/ab-backend/internal/modules/customers/application"
	"github.com/naelcodes/ab-backend/internal/pkg/server"
)

type CostumerRestController struct {
	Application *application.CustomerApplication
}

func (controller *CostumerRestController) Init(appEngine *server.AppEngine) {

	//TODO : Replace this with global config
	baseRouter := appEngine.Get().Group("/api/v1")

	//validation middlewares here
	controller.attachCustomerRoutesHandlers(baseRouter.Group("/customers"))
}

func (controller *CostumerRestController) attachCustomerRoutesHandlers(router fiber.Router) {

	router.Get("", controller.GetAllCustomersHandler())
	router.Get("/:id", controller.GetCustomerHandler())
	router.Post("", controller.CreateCustomerHandler())
	router.Put("/:id", controller.UpdateCustomerHandler())
	router.Delete("/:id", controller.DeleteCustomerHandler())
}
