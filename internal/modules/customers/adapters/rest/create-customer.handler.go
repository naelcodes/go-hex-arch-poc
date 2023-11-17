package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naelcodes/ab-backend/internal/modules/customers/application"
)

func (cm *CostumerModule) CreateCustomerHandler() fiber.Handler {

	appSvc := new(application.CreateCustomerAppSvc)
	appSvc.Init(cm.Repository)

	return func(c *fiber.Ctx) error {

		result, err := appSvc.Execute()

		if err != nil {

			return c.Status(500).JSON(err)
		}

		return c.Status(200).JSON(result)

	}

}
