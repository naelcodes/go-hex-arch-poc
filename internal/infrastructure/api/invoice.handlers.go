package api

import "github.com/gofiber/fiber/v2"

func (controller *RestController) CreateInvoiceHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON("\\Create Invoice handler")

	}
}

func (controller *RestController) GetAllInvoiceHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON("/getAllInvoice handler")

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
