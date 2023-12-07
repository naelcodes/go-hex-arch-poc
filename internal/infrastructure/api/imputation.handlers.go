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

		err = controller.ApplicationService.ApplyInvoiceImputationService(types.EID(id), InvoiceImputationDTOList)

		if err != nil {
			utils.Logger.Error(fmt.Sprintf("[ApplyInvoiceImputationHandler] - Error applying invoice imputation: %v", err))
			return err
		}

		utils.Logger.Info(fmt.Sprintf("[ApplyInvoiceImputationHandler] - Applied invoice imputation"))

		return c.Status(200).JSON(fiber.Map{
			"success": true,
		})

	}
}

func (controller *RestController) GetInvoiceImputationsHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON("/getAllInvoice handler")

	}
}
