package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	"github.com/naelcodes/ab-backend/pkg/errors"
	"github.com/naelcodes/ab-backend/pkg/types"
	"github.com/naelcodes/ab-backend/pkg/utils"
)

func (controller *RestController) ApplyInvoiceImputationHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {

		utils.Logger.Info("[ApplyInvoiceImputationHandler] - Applying invoice imputation")

		id, err := c.ParamsInt("id")

		if err != nil {
			utils.Logger.Error(fmt.Sprintf("[ApplyInvoiceImputationHandler] - Error parsing invoice id: %v", err))
			return errors.ServiceError(err, "Parsing invoice id")
		}

		utils.Logger.Info(fmt.Sprintf("[ApplyInvoiceImputationHandler] - Invoice id: %v", id))

		InvoiceImputationDTOList := c.Locals("payload").([]*dto.InvoiceImputationDTO)

		utils.Logger.Info(fmt.Sprintf("[ApplyInvoiceImputationHandler] - payload : %v", InvoiceImputationDTOList))

		imputationOperationResult, err := controller.ApplicationService.ApplyInvoiceImputationService(types.EID(id), InvoiceImputationDTOList)

		if err != nil {
			utils.Logger.Error(fmt.Sprintf("[ApplyInvoiceImputationHandler] - Error applying invoice imputation: %v", err))
			return err
		}

		utils.Logger.Info("[ApplyInvoiceImputationHandler] - Applied invoice imputation")

		return c.Status(fiber.StatusOK).JSON(imputationOperationResult)

	}
}

func (controller *RestController) GetInvoiceImputationsHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {

		utils.Logger.Info("[GetInvoiceImputationsHandler] - Getting invoice imputations")

		id, err := c.ParamsInt("id")

		if err != nil {
			utils.Logger.Error(fmt.Sprintf("[GetInvoiceImputationsHandler] - Error parsing invoice id: %v", err))
			return errors.ServiceError(err, "Parsing invoice id")
		}

		utils.Logger.Info(fmt.Sprintf("[GetInvoiceImputationsHandler] - Invoice id: %v", id))

		InvoiceImputationDTOList, err := controller.ApplicationService.GetInvoiceImputationService(types.EID(id))

		if err != nil {
			utils.Logger.Error(fmt.Sprintf("[GetInvoiceImputationsHandler] - Error getting invoice imputations: %v", err))
			return err
		}

		utils.Logger.Info(fmt.Sprintf("[GetInvoiceImputationsHandler] - Invoice imputations: %v", InvoiceImputationDTOList))

		return c.Status(fiber.StatusOK).JSON(InvoiceImputationDTOList)

	}
}
