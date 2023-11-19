package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/naelcodes/ab-backend/internal/pkg/middleware"
)

type AppEngine struct {
	server *fiber.App
}

func (appEngine *AppEngine) Init() {

	app := fiber.New()
	app.Use(middleware.Cors())

	appEngine.server = app
}

func (appEngine *AppEngine) Serve() error {
	//TODO : Update this to environment config
	fmt.Println("Server Running on port 3000....")
	return appEngine.server.Listen(":3000")
}

func (appEngine *AppEngine) Get() *fiber.App {
	return appEngine.server
}
