package api

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	"github.com/naelcodes/ab-backend/pkg/errors"
	"github.com/naelcodes/ab-backend/pkg/types"
	"github.com/naelcodes/ab-backend/pkg/utils"
)

func (controller *RestController) CreatePaymentHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {

		utils.Logger.Info("[CreatePaymentHandler] - Creating payment")

		createPaymentDto := c.Locals("payload").(*dto.CreatePaymentDTO)

		utils.Logger.Info(fmt.Sprintf("[CreatePaymentHandler] - Payload: %v", createPaymentDto))

		newPaymentDTO, err := controller.ApplicationService.CreatePaymentService(createPaymentDto)

		if err != nil {
			utils.Logger.Error(fmt.Sprintf("[CreatePaymentHandler] - Error creating payment: %v", err))
			return err
		}

		utils.Logger.Info(fmt.Sprintf("[CreatePaymentHandler] - Created payment DTO: %v", newPaymentDTO))
		return c.Status(fiber.StatusCreated).JSON(newPaymentDTO)

	}
}

func (controller *RestController) GetAllPaymentsHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {

		utils.Logger.Info("[GetAllPaymentsHandler] - Getting all payments")

		queryParams := new(types.GetQueryParams)
		err := c.QueryParser(queryParams)

		utils.Logger.Info(fmt.Sprintf("[GetAllPaymentsHandler] - Query params: %v", queryParams))

		if err != nil {
			utils.Logger.Error(fmt.Sprintf("[GetAllPaymentsHandler] - Error parsing query params: %v", err))
			return errors.ServiceError(err, "Parsing query params")
		}

		getAllPaymentsDTO, err := controller.ApplicationService.GetAllPaymentsService(queryParams)

		if err != nil {
			utils.Logger.Error(fmt.Sprintf("[GetAllPaymentsHandler] - Error getting all payments: %v", err))
			return err
		}

		utils.Logger.Info(fmt.Sprintf("[GetAllPaymentsHandler] - Get all payments DTO: %v", getAllPaymentsDTO))

		return c.Status(fiber.StatusOK).JSON(getAllPaymentsDTO)

	}
}

func (controller *RestController) GetPaymentHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {

		utils.Logger.Info("[GetPaymentHandler] - Getting payment")
		id, err := c.ParamsInt("id")
		if err != nil {
			utils.Logger.Error(fmt.Sprintf("[GetPaymentHandler] - Error parsing id: %v", err))
			return errors.ServiceError(err, "Id Parsing in URL parameter")
		}

		utils.Logger.Info(fmt.Sprintf("[GetPaymentHandler] - Id: %d", id))

		paymentDTO, err := controller.ApplicationService.GetPaymentService(types.EID(id))
		if err != nil {
			utils.Logger.Error(fmt.Sprintf("[GetPaymentHandler] - Error getting payment: %v", err))
			return err
		}

		utils.Logger.Info(fmt.Sprintf("[GetPaymentHandler] - Payment DTO: %v", paymentDTO))

		return c.Status(fiber.StatusOK).JSON(paymentDTO)

	}
}

func (controller *RestController) UpdatePaymentHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON("/updatePayment handler")

	}
}

func (controller *RestController) DeletePaymentHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON("/deletePayment handler")

	}
}
