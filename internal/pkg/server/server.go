package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/naelcodes/ab-backend/internal/pkg/middleware"
)

type Engine struct {
	server *fiber.App
}

func (e *Engine) Init() {

	app := fiber.New()
	app.Use(middleware.Cors())

	e.server = app
}

func (e *Engine) Serve() error {
	//TODO : Update this to environment config
	fmt.Println("Server Running on port 3000....")
	return e.server.Listen(":3000")
}

func (e *Engine) Get() *fiber.App {
	return e.server
}
