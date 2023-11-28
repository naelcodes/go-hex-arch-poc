package api

import "github.com/gofiber/fiber/v2"

func (controller *RestController) CreatePaymentHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON("\\Create Payment handler")

	}
}

func (controller *RestController) GetAllPaymentsHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON("/getAllPayment handler")

	}
}

func (controller *RestController) GetPaymentHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON("/getPayment handler")

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
