package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	"github.com/naelcodes/ab-backend/pkg/errors"
	"github.com/naelcodes/ab-backend/pkg/types"
	"github.com/naelcodes/ab-backend/pkg/utils"
)

func (controller *RestController) CreateInvoiceHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {

		utils.Logger.Info("[CreateInvoiceHandler] - Creating invoice")

		createInvoiceDto := c.Locals("payload").(*dto.CreateInvoiceDTO)

		utils.Logger.Info(fmt.Sprintf("[CreateInvoiceHandler] - Payload: %v", createInvoiceDto))

		newInvoiceDTO, err := controller.ApplicationService.CreateInvoiceService(createInvoiceDto)

		if err != nil {
			utils.Logger.Error(fmt.Sprintf("[CreateInvoiceHandler] - Error creating invoice: %v", err))
			return err
		}

		utils.Logger.Info(fmt.Sprintf("[CreateInvoiceHandler] - Created invoice DTO: %v", newInvoiceDTO))
		return c.Status(fiber.StatusOK).JSON(newInvoiceDTO)

	}
}

func (controller *RestController) GetAllInvoiceHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		queryParams := new(types.GetQueryParams)
		err := c.QueryParser(queryParams)

		utils.Logger.Info(fmt.Sprintf("[GetAllInvoiceHandler] - Query params: %v", queryParams))

		if err != nil {
			utils.Logger.Error(fmt.Sprintf("[GetAllInvoiceHandler] - Error parsing query params: %v", err))
			return errors.ServiceError(err, "Parsing query params")
		}

		getAllInvoiceDTO, err := controller.ApplicationService.GetAllInvoiceService(queryParams)

		if err != nil {
			utils.Logger.Error(fmt.Sprintf("[GetAllInvoiceHandler] - Error getting all invoice DTO: %v", err))
			return err
		}

		utils.Logger.Info(fmt.Sprintf("[GetAllInvoiceHandler] - All invoice DTO: %v", getAllInvoiceDTO))
		return c.Status(fiber.StatusOK).JSON(getAllInvoiceDTO)

	}
}

func (controller *RestController) GetInvoiceHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {

		id, err := c.ParamsInt("id")
		if err != nil {
			utils.Logger.Error(fmt.Sprintf("[GetInvoiceHandler] - Error parsing id: %v", err))
			return errors.ServiceError(err, "Id Parsing in URL parameter")
		}
		utils.Logger.Info(fmt.Sprintf("[GetInvoiceHandler] - params Id: %v", id))

		queryParams := new(types.GetQueryParams)
		err = c.QueryParser(queryParams)

		if err != nil {
			utils.Logger.Error(fmt.Sprintf("[GetInvoiceHandler] - Error parsing query params: %v", err))
			return errors.ServiceError(err, "Parsing query params")
		}

		getInvoiceDTO, err := controller.ApplicationService.GetInvoiceService(id, queryParams)

		if err != nil {
			utils.Logger.Error(fmt.Sprintf("[GetInvoiceHandler] - Error getting invoice DTO: %v", err))
			return err
		}

		utils.Logger.Info(fmt.Sprintf("[GetInvoiceHandler] - Invoice DTO: %v", getInvoiceDTO))
		return c.Status(fiber.StatusOK).JSON(getInvoiceDTO)

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
