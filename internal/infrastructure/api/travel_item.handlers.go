package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naelcodes/ab-backend/pkg/errors"
	"github.com/naelcodes/ab-backend/pkg/types"
)

func (controller *RestController) GetAllTravelItemsHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {

		queryParams := new(types.GetQueryParams)
		err := c.QueryParser(queryParams)

		if err != nil {
			return errors.ServiceError(err, "Parsing query params")
		}

		getAllTravelItemsDTO, err := controller.ApplicationService.GetAllTravelItemService(queryParams)

		if err != nil {
			return err
		}

		return c.Status(fiber.StatusOK).JSON(getAllTravelItemsDTO)

	}

}
