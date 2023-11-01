package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/naelcodes/customer-rest-api/internal/database"
	"github.com/naelcodes/customer-rest-api/internal/router"
)



func main(){

	 database.ConnectDb()

	 app := fiber.New()
	 app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
	}))

	 apiV1Router := app.Group("/api/v1")
	 customerRouter := apiV1Router.Group("/customers")

	 router.SetupRoutes(customerRouter)

	 fmt.Println("SERVER STARTING ON PORT 3000....")
	 log.Fatal(app.Listen(":3000"))
	 
}