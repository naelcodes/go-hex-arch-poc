package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/naelcodes/ab-backend/internal/pkg/logger"
	"github.com/naelcodes/ab-backend/internal/pkg/middleware"
	"github.com/naelcodes/ab-backend/internal/pkg/pubsub"
)

type AppEngine struct {
	server *fiber.App
	broker *pubsub.Broker
	logger *logger.Logger
}

func (appEngine *AppEngine) Init() {

	app := fiber.New()
	app.Use(middleware.Cors())

	appEngine.server = app

	appEngine.logger = logger.NewLogger()
	appEngine.broker = pubsub.NewBroker()

}

func (appEngine *AppEngine) Serve() error {
	appEngine.logger.Info("Server Running on port 3000....")
	return appEngine.server.Listen(":3000")
}

func (appEngine *AppEngine) Get() *fiber.App {
	return appEngine.server
}

func (appEngine *AppEngine) GetBroker() *pubsub.Broker {
	return appEngine.broker
}

func (appEngine *AppEngine) GetLogger() *logger.Logger {
	return appEngine.logger
}
