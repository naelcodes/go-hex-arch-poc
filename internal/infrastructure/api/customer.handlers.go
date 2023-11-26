package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	"github.com/naelcodes/ab-backend/pkg/types"
)

func (controller *RestController) GetAllCustomersHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {

		queryParams := new(types.GetQueryParams)
		c.QueryParser(queryParams)

		result, err := controller.ApplicationService.GetAllCustomersService(queryParams)

		if err != nil {
			fmt.Println("error", err)
			return c.Status(500).JSON(err)
		}
		return c.Status(200).JSON(result)

	}

}

func (controller *RestController) GetCustomerHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON("/getAllCustomerHandler")

	}
}

func (controller *RestController) CreateCustomerHandler() fiber.Handler {

	return func(c *fiber.Ctx) error {
		customer_dto := c.Locals("customer_dto").(*dto.CreateCustomerDTO)
		id, err := controller.ApplicationService.CreateCustomerService(customer_dto)
		if err != nil {
			return c.Status(500).JSON(err)
		}
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"id": id,
		})
	}

}

func (controller *RestController) UpdateCustomerHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON("/getAllCustomerHandler")

	}
}

func (controller *RestController) DeleteCustomerHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON("\\delete handler")

	}

}
