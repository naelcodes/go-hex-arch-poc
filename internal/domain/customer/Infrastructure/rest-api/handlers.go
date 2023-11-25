package restAdapter

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/naelcodes/ab-backend/internal/common"
	"github.com/naelcodes/ab-backend/internal/modules/customers/dto"
)

func (controller *CostumerRestController) GetAllCustomersHandler() fiber.Handler {

	return func(c *fiber.Ctx) error {

		queryParams := new(common.GetQueryParams)
		c.QueryParser(queryParams)

		result, err := controller.Application.GetAllCustomersService(queryParams)

		if err != nil {
			fmt.Println("error", err)
			return c.Status(500).JSON(err)
		}
		return c.Status(200).JSON(result)

	}

}

func (controller *CostumerRestController) GetCustomerHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON("/getAllCustomerHandler")

	}
}

func (controller *CostumerRestController) CreateCustomerHandler() fiber.Handler {

	return func(c *fiber.Ctx) error {
		customer_dto := c.Locals("customer_dto").(*dto.CreateCustomerDTO)
		id, err := controller.Application.CreateCustomerService(customer_dto)
		if err != nil {
			return c.Status(500).JSON(err)
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"id": id,
		})
	}

}

func (controller *CostumerRestController) UpdateCustomerHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON("/getAllCustomerHandler")

	}
}

func (controller *CostumerRestController) DeleteCustomerHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON("\\delete handler")

	}

}
