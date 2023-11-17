package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naelcodes/ab-backend/internal/modules/customers/domain"
	"github.com/naelcodes/ab-backend/internal/pkg/server"
)

type CostumerModule struct {
	Repository domain.ICustomerRepository
}

func (cm *CostumerModule) Init(e *server.Engine, repository domain.ICustomerRepository) {

	//TODO : Replace this with global config
	apiGroup := e.Get().Group("/api/v1")

	//validation middlewares here

	cm.Repository = repository

	cm.attachCustomerHandlers(apiGroup.Group("/customers"))
	cm.attachCountriesHandler(apiGroup.Group("/countries"))
}

func (cm *CostumerModule) attachCustomerHandlers(router fiber.Router) {

	// router.Get("", c.GetAlLCustomers)
	// router.Get("/:id", c.GetAlLCustomers)
	// router.Post("", c.CreateCustomer)
	// router.Put("/:id", c.UpdateCustomer)
	// router.Delete("/:id", c.DeleteCustomer)
}

func (cm *CostumerModule) attachCountriesHandler(router fiber.Router) {
	router.Get("", cm.GetAllCountriesHandler())
}
