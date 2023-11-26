package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naelcodes/ab-backend/internal/modules/customers/dto"
)

func CustomerValidationMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		switch c.Method() {

		case fiber.MethodGet:
			return c.Next()

		case fiber.MethodPost:
			payload := new(dto.CreateCustomerDTO)
			if err := c.BodyParser(payload); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
					"success": false,
					"error":   err.Error(),
				})
			}

			//validate payload
			if err := payload.Validate(); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
					"success": false,
					"error":   err.Error(),
				})
			}

			// store validate DTO in the context
			c.Locals("customer_dto", payload)
			return c.Next()

		case fiber.MethodPatch:
			payload := new(dto.UpdateCustomerDTO)
			if err := c.BodyParser(payload); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
					"success": false,
					"error":   err.Error(),
				})
			}

			//validate payload
			if err := payload.Validate(); err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(&fiber.Map{
					"success": false,
					"error":   err.Error(),
				})
			}

			// store validate DTO in the context
			c.Locals("customer_dto", payload)
			return c.Next()
			
		case fiber.MethodDelete:
			return c.Next()
		}
		return c.Next()
	}
}
