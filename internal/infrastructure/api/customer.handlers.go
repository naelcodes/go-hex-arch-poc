package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	"github.com/naelcodes/ab-backend/pkg/errors"
	"github.com/naelcodes/ab-backend/pkg/types"
	"github.com/naelcodes/ab-backend/pkg/utils"
)

func (controller *RestController) GetAllCustomersHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {

		queryParams := new(types.GetQueryParams)
		err := c.QueryParser(queryParams)

		utils.Logger.Info(fmt.Sprintf("[GetAllCustomersHandler] - Query params: %v", queryParams))

		if err != nil {
			utils.Logger.Error(fmt.Sprintf("[GetAllCustomersHandler] - Error parsing query params: %v", err))
			return errors.ServiceError(err, "Parsing query params")
		}

		getAllCustomersDTO, err := controller.ApplicationService.GetAllCustomersService(queryParams)

		if err != nil {
			utils.Logger.Error(fmt.Sprintf("[GetAllCustomersHandler] - Error getting all customers DTO: %v", err))
			return err
		}

		utils.Logger.Info(fmt.Sprintf("[GetAllCustomersHandler] - All customers DTO: %v", getAllCustomersDTO))
		return c.Status(fiber.StatusOK).JSON(getAllCustomersDTO)

	}

}

func (controller *RestController) GetCustomerHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")

		if err != nil {
			utils.Logger.Error(fmt.Sprintf("[GetCustomerHandler] - Error parsing id: %v", err))
			return errors.ServiceError(err, "Id Parsing in URL parameter")
		}

		utils.Logger.Info(fmt.Sprintf("[GetCustomerHandler] - params Id: %v", id))

		customerDTO, err := controller.ApplicationService.GetCustomerService(types.EID(id))

		if err != nil {
			utils.Logger.Error(fmt.Sprintf("[GetCustomerHandler] - Error getting customer DTO: %v", err))
			return err
		}

		utils.Logger.Info(fmt.Sprintf("[GetCustomerHandler] - Customer DTO: %v", customerDTO))

		return c.Status(fiber.StatusOK).JSON(customerDTO)

	}
}

func (controller *RestController) CreateCustomerHandler() fiber.Handler {

	return func(c *fiber.Ctx) error {
		createCustomerDto := c.Locals("payload").(*dto.CreateCustomerDTO)

		utils.Logger.Info(fmt.Sprintf("[CreateCustomerHandler] - Payload: %v", createCustomerDto))

		newCustomerDTO, err := controller.ApplicationService.CreateCustomerService(createCustomerDto)

		if err != nil {
			utils.Logger.Error(fmt.Sprintf("[CreateCustomerHandler] - Error creating customer: %v", err))
			return err
		}

		utils.Logger.Info(fmt.Sprintf("[CreateCustomerHandler] - Created customer DTO: %v", newCustomerDTO))

		return c.Status(fiber.StatusCreated).JSON(newCustomerDTO)
	}

}

func (controller *RestController) UpdateCustomerHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {

		updateCustomerDto := c.Locals("payload").(*dto.UpdateCustomerDTO)

		utils.Logger.Info(fmt.Sprintf("[UpdateCustomerHandler] - Payload: %v", updateCustomerDto))

		id, err := c.ParamsInt("id")

		if err != nil {
			utils.Logger.Error(fmt.Sprintf("[UpdateCustomerHandler] - Error parsing id: %v", err))
			return errors.ServiceError(err, "Id Parsing in URL parameter")
		}

		utils.Logger.Info(fmt.Sprintf("[UpdateCustomerHandler] - params Id: %v", id))

		RecordWasUpdated, err := controller.ApplicationService.UpdateCustomerService(id, updateCustomerDto)

		if err != nil {
			utils.Logger.Error(fmt.Sprintf("[UpdateCustomerHandler] - Error updating customer: %v", err))
			return err
		}

		utils.Logger.Info(fmt.Sprintf("[UpdateCustomerHandler] - Updated customer DTO: %v", RecordWasUpdated))

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": RecordWasUpdated,
		})

	}
}

func (controller *RestController) DeleteCustomerHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {

		id, err := c.ParamsInt("id")
		if err != nil {
			utils.Logger.Error(fmt.Sprintf("[DeleteCustomerHandler] - Error parsing id: %v", err))
			return errors.ServiceError(err, "Id Parsing in URL parameter")
		}

		utils.Logger.Info(fmt.Sprintf("[DeleteCustomerHandler] - params Id: %v", id))

		RecordWasDeleted, err := controller.ApplicationService.DeleteCustomerService(types.EID(id))

		if err != nil {
			utils.Logger.Error(fmt.Sprintf("[DeleteCustomerHandler] - Error deleting customer: %v", err))
			return err
		}

		utils.Logger.Info(fmt.Sprintf("[DeleteCustomerHandler] - Deleted customer DTO: %v", RecordWasDeleted))

		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": RecordWasDeleted,
		})

	}

}

func (controller *RestController) GetCustomerPaymentsHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			utils.Logger.Error(fmt.Sprintf("[GetCustomerPaymentsHandler] - Error parsing id: %v", err))
			return errors.ServiceError(err, "Id Parsing in URL parameter")
		}

		utils.Logger.Info(fmt.Sprintf("[GetCustomerPaymentsHandler] - params Id: %v", id))

		queryParams := new(types.GetQueryParams)
		err = c.QueryParser(queryParams)

		if err != nil {
			utils.Logger.Error(fmt.Sprintf("[GetCustomerPaymentsHandler] - Error parsing query params: %v", err))
			return errors.ServiceError(err, "Parsing query params")
		}

		utils.Logger.Info(fmt.Sprintf("[GetCustomerPaymentsHandler] - Query params: %v", queryParams))

		isOpen := c.QueryBool("open", false)

		utils.Logger.Info(fmt.Sprintf("[GetCustomerPaymentsHandler] - Query params: %v", isOpen))

		getAllCustomerPaymentsDTO, err := controller.ApplicationService.GetCustomerPaymentsService(types.EID(id), queryParams, &isOpen)

		if err != nil {
			utils.Logger.Error(fmt.Sprintf("[GetCustomerPaymentsHandler] - Error getting all customer payments DTO: %v", err))
			return err
		}

		utils.Logger.Info(fmt.Sprintf("[GetCustomerPaymentsHandler] - All customer payments DTO: %v", getAllCustomerPaymentsDTO))

		return c.Status(fiber.StatusOK).JSON(getAllCustomerPaymentsDTO)

	}
}

func (controller *RestController) GetCustomerInvoicesHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {

		id, err := c.ParamsInt("id")
		if err != nil {
			utils.Logger.Error(fmt.Sprintf("[GetCustomerInvoicesHandler] - Error parsing id: %v", err))
			return errors.ServiceError(err, "Id Parsing in URL parameter")
		}

		utils.Logger.Info(fmt.Sprintf("[GetCustomerInvoicesHandler] - params Id: %v", id))

		queryParams := new(types.GetQueryParams)
		err = c.QueryParser(queryParams)

		if err != nil {
			utils.Logger.Error(fmt.Sprintf("[GetCustomerInvoicesHandler] - Error parsing query params: %v", err))
			return errors.ServiceError(err, "Parsing query params")
		}

		utils.Logger.Info(fmt.Sprintf("[GetCustomerInvoicesHandler] - Query params: %v", queryParams))

		paid := c.QueryBool("paid", false)

		utils.Logger.Info(fmt.Sprintf("[GetCustomerInvoicesHandler] - Query params: %v", paid))

		getAllCustomerInvoicesDTO, err := controller.ApplicationService.GetCustomerInvoicesService(types.EID(id), queryParams, &paid)

		if err != nil {
			utils.Logger.Error(fmt.Sprintf("[GetCustomerInvoicesHandler] - Error getting all customer invoices DTO: %v", err))
			return err
		}

		utils.Logger.Info(fmt.Sprintf("[GetCustomerInvoicesHandler] - All customer invoices DTO: %v", getAllCustomerInvoicesDTO))

		return c.Status(fiber.StatusOK).JSON(getAllCustomerInvoicesDTO)
	}

}
