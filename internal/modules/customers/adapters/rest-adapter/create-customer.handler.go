package restAdapter

import (
	"github.com/gofiber/fiber/v2"
)

func (controller *CostumerRestController) CreateCustomerHandler() fiber.Handler {

	return func(c *fiber.Ctx) error {
		// customerDTO := new(acl.CreateCustomerDTO)
		// if err := c.BodyParser(customerDTO); err != nil {
		// 	return c.Status(400).JSON(err)
		// }

		// result, err := customerRestAdapter.Application.CreateCustomerService()

		// if err != nil {
		// 	return c.Status(500).JSON(err)
		// }

		return c.Status(201).JSON("")

	}

}
