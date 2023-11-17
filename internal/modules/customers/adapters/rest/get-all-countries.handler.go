package rest

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naelcodes/ab-backend/internal/modules/customers/application/services"
)

func (cm *CostumerModule) GetAllCountriesHandler() fiber.Handler {

	appSvc := new(services.GetAllCountriesAppSvc)
	appSvc.Init(cm.Repository)

	return func(c *fiber.Ctx) error {

		result, err := appSvc.Execute()

		if err != nil {

			return c.Status(500).JSON(err)
		}

		return c.Status(200).JSON(result)

	}

}
