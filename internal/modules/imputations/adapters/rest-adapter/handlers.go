package restAdapter

import "github.com/gofiber/fiber/v2"

func (controller *ImputationRestController) GetInvoiceImputationsHandler() fiber.Handler {

	return func(c *fiber.Ctx) error {
		return c.Status(201).JSON("")

	}

}

func (controller *ImputationRestController) UpdateInvoiceImputationsHandler() fiber.Handler {

	return func(c *fiber.Ctx) error {
		return c.Status(201).JSON("")

	}

}

func (controller *ImputationRestController) ApplyImputationsHandler() fiber.Handler {

	return func(c *fiber.Ctx) error {
		return c.Status(201).JSON("")

	}

}
