package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naelcodes/ab-backend/handlers"
)

func SetupRoutes(customerRouter fiber.Router) {

	//TODO : add paginations on GET
	customerRouter.Get("", handlers.GetAlLCustomers)
	customerRouter.Post("", handlers.AddCustomer)
	customerRouter.Put("/:id", handlers.UpdateCustomer)
	customerRouter.Delete("/:id", handlers.DeleteCustomer)
}
