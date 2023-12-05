package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	"github.com/naelcodes/ab-backend/pkg/errors"
	"github.com/naelcodes/ab-backend/pkg/types"
)

func (controller *RestController) CreatePaymentHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		createPaymentDto := c.Locals("payload").(*dto.CreatePaymentDTO)
		newPaymentDTO, err := controller.ApplicationService.CreatePaymentService(createPaymentDto)
		if err != nil {
			return err
		}
		return c.Status(fiber.StatusCreated).JSON(newPaymentDTO)

	}
}

func (controller *RestController) GetAllPaymentsHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		queryParams := new(types.GetQueryParams)
		err := c.QueryParser(queryParams)

		if err != nil {
			return errors.ServiceError(err, "Parsing query params")
		}

		getAllPaymentsDTO, err := controller.ApplicationService.GetAllPaymentsService(queryParams)

		if err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(getAllPaymentsDTO)

	}
}

func (controller *RestController) GetPaymentHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return errors.ServiceError(err, "Id Parsing in URL parameter")
		}
		paymentDTO, err := controller.ApplicationService.GetPaymentService(types.EID(id))
		if err != nil {
			return err
		}

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
