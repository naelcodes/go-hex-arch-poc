package api

import "github.com/gofiber/fiber/v2"

func (controller *RestController) GetAllTravelItemsHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(200).JSON("GetAll TravelITems")

	}

}
