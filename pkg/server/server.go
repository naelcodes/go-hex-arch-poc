package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"

	"github.com/naelcodes/ab-backend/config"
	"github.com/naelcodes/ab-backend/pkg/errors"
	"github.com/naelcodes/ab-backend/pkg/utils"
)

type AppEngine struct {
	server *fiber.App
	logger *utils.ZeroLogger
}

func (appEngine *AppEngine) Init() {

	app := fiber.New(fiber.Config{
		ErrorHandler: errors.GlobalErrorHandler,
	})
	appEngine.server = app
	appEngine.logger = utils.NewZeroLogger()

}

func (appEngine *AppEngine) Start() error {
	appEngine.logger.Info(fmt.Sprintf("Server Running on port :%v....", config.APP_ENGINE_SERVER_PORT))
	return appEngine.server.Listen(fmt.Sprintf(":%v", config.APP_ENGINE_SERVER_PORT))
}

func (appEngine *AppEngine) GetServer() *fiber.App {
	return appEngine.server
}

func (appEngine *AppEngine) GetLogger() *utils.ZeroLogger {
	return appEngine.logger
}
