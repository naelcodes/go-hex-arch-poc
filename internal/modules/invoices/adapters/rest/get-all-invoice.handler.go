package rest

import "github.com/gofiber/fiber/v2"

func (im *InvoiceModule) GetAllInvoiceHandler() fiber.Handler {

	//appSvc := new(application.GetAllCountriesService)
	//appSvc.Init(cm.Repository)

	return func(c *fiber.Ctx) error {

		//result, err := appSvc.Execute()

		// if err != nil {
		// 	return c.Status(500).JSON(err)
		// }

		return c.Status(200).JSON("\\get all handler")

	}

}
