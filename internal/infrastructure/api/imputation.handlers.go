package api

import (
	"github.com/gofiber/fiber/v2"
)

func (controller *RestController) AddInvoiceImputationHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON("\\Create Invoice handler")

	}
}

func (controller *RestController) GetInvoiceImputationsHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON("/getAllInvoice handler")

	}
}

func (controller *RestController) UpdateInvoiceImputationsHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON("/updateInvoice handler")

	}
}
