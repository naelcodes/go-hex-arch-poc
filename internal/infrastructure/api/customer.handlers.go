package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	"github.com/naelcodes/ab-backend/pkg/errors"
	"github.com/naelcodes/ab-backend/pkg/types"
)

func (controller *RestController) GetAllCustomersHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {

		queryParams := new(types.GetQueryParams)
		err := c.QueryParser(queryParams)

		if err != nil {
			return errors.ServiceError(err, "Parsing query params")
		}

		getAllCustomersDTO, err := controller.ApplicationService.GetAllCustomersService(queryParams)

		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(getAllCustomersDTO)

	}

}

func (controller *RestController) GetCustomerHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return errors.ServiceError(err, "Id Parsing in URL parameter")
		}
		customerDTO, err := controller.ApplicationService.GetCustomerService(types.EID(id))
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(customerDTO)

	}
}

func (controller *RestController) CreateCustomerHandler() fiber.Handler {

	return func(c *fiber.Ctx) error {
		createCustomerDto := c.Locals("payload").(*dto.CreateCustomerDTO)
		newCustomerDTO, err := controller.ApplicationService.CreateCustomerService(createCustomerDto)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusCreated).JSON(newCustomerDTO)
	}

}

func (controller *RestController) UpdateCustomerHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		updateCustomerDto := c.Locals("payload").(*dto.UpdateCustomerDTO)
		RecordWasUpdated, err := controller.ApplicationService.UpdateCustomerService(updateCustomerDto)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": RecordWasUpdated,
		})

	}
}

func (controller *RestController) DeleteCustomerHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {

		id, err := c.ParamsInt("id")
		if err != nil {
			return errors.ServiceError(err, "Id Parsing in URL parameter")
		}
		RecordWasDeleted, err := controller.ApplicationService.DeleteCustomerService(types.EID(id))
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": RecordWasDeleted,
		})

	}

}
