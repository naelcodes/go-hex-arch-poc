package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	"github.com/naelcodes/ab-backend/pkg/errors"
	"github.com/naelcodes/ab-backend/pkg/types"
)

func (controller *RestController) CreateInvoiceHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {

		controller.Logger.Info("[CreateInvoiceHandler] - Creating invoice")

		createInvoiceDto := c.Locals("payload").(*dto.CreateInvoiceDTO)

		controller.Logger.Info(fmt.Sprintf("[CreateInvoiceHandler] - Payload: %v", createInvoiceDto))

		newInvoiceDTO, err := controller.ApplicationService.CreateInvoiceService(createInvoiceDto)

		if err != nil {
			controller.Logger.Error(fmt.Sprintf("[CreateInvoiceHandler] - Error creating invoice: %v", err))
			return err
		}

		controller.Logger.Info(fmt.Sprintf("[CreateInvoiceHandler] - Created invoice DTO: %v", newInvoiceDTO))
		return c.Status(fiber.StatusOK).JSON(newInvoiceDTO)

	}
}

func (controller *RestController) GetAllInvoiceHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		queryParams := new(types.GetQueryParams)
		err := c.QueryParser(queryParams)

		controller.Logger.Info(fmt.Sprintf("[GetAllInvoiceHandler] - Query params: %v", queryParams))

		if err != nil {
			controller.Logger.Error(fmt.Sprintf("[GetAllInvoiceHandler] - Error parsing query params: %v", err))
			return errors.ServiceError(err, "Parsing query params")
		}

		getAllInvoiceDTO, err := controller.ApplicationService.GetAllInvoiceService(queryParams)

		if err != nil {
			controller.Logger.Error(fmt.Sprintf("[GetAllInvoiceHandler] - Error getting all invoice DTO: %v", err))
			return err
		}

		controller.Logger.Info(fmt.Sprintf("[GetAllInvoiceHandler] - All invoice DTO: %v", getAllInvoiceDTO))
		return c.Status(fiber.StatusOK).JSON(getAllInvoiceDTO)

	}
}

func (controller *RestController) GetInvoiceHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON("/getInvoice handler")

	}
}

func (controller *RestController) UpdateInvoiceHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON("/updateInvoice handler")

	}
}

func (controller *RestController) DeleteInvoiceHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON("/deleteInvoice handler")

	}
}
