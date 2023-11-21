package restAdapter

import "github.com/gofiber/fiber/v2"

func (controller *InvoiceRestController) GetAllInvoiceHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON("\\createInvoice handler")

	}

}

func (controller *InvoiceRestController) GetInvoiceHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON("\\createInvoice handler")

	}

}
func (controller *InvoiceRestController) CreateInvoiceHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON("\\createInvoice handler")

	}

}

func (controller *InvoiceRestController) UpdateInvoiceHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON("\\createInvoice handler")

	}

}

func (controller *InvoiceRestController) DeleteInvoiceHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON("\\createInvoice handler")

	}

}

func (controller *InvoiceRestController) GetAllTravelItemsHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON("\\createInvoice handler")

	}

}
