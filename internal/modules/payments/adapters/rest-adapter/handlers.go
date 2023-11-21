package restAdapter

import "github.com/gofiber/fiber/v2"

func (controller *PaymentRestController) AddCustomerPaymentHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON("\\createInvoice handler")

	}

}

func (controller *PaymentRestController) GetCustomerPaymentHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON("\\createInvoice handler")

	}

}

func (controller *PaymentRestController) GetAllPaymentHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON("\\createInvoice handler")

	}

}

func (controller *PaymentRestController) UpdatePaymentHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON("\\createInvoice handler")

	}

}

func (controller *PaymentRestController) DeletePaymentHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON("\\createInvoice handler")

	}

}
