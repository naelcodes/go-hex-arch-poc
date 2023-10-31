package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naelcodes/customer-rest-api/internal/database"
	"github.com/naelcodes/customer-rest-api/pkg/models"
)

func AddCustomer(c *fiber.Ctx) error {
	newCustomer := new(models.Customer)

		if err := c.BodyParser(newCustomer); err != nil {
			return err
		}

		result := database.DB.Create(&newCustomer)
		if result.Error != nil {
			return c.Status(500).JSON(result.Error)
		}

		return c.Status(201).JSON(fiber.Map{
			"userID":newCustomer.Id,
		})
}

func UpdateCustomer(c *fiber.Ctx) error {
	newCustomer := new(models.Customer)
		if err := c.BodyParser(newCustomer); err != nil {
			return err
		}

		id, err := c.ParamsInt("id");
		if err != nil{
			return err
		}
		
		result := database.DB.Model(&newCustomer).Where("id= ?",id).Updates(newCustomer)
		if result.Error != nil {
			return c.Status(500).JSON(result.Error)
		}

		return c.Status(201).JSON(fiber.Map{
			"updateCount":result.RowsAffected,
		})
}

func GetAlLCustomers(c *fiber.Ctx) error {
		customers := []models.Customer{}
		database.DB.Find(&customers)
		c.SendStatus(200)	
		return c.Status(200).JSON(customers)	
}


func DeleteCustomer(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id");
		if err != nil{
			return err
		}

		result := database.DB.Delete(&models.Customer{},id)
		if result.Error != nil {
			return c.Status(500).JSON(result.Error)
		}

		if result.RowsAffected == 0{
			return  c.Status(404).JSON(fiber.Map{
				"DeletedCount":result.RowsAffected,
			})
		}

		return c.Status(201).JSON(fiber.Map{
			"DeletedCount":result.RowsAffected,
		})
}