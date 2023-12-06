package main

import (
	"context"
	"log"

	"github.com/naelcodes/ab-backend/config"
	"github.com/naelcodes/ab-backend/config/database"
	"github.com/naelcodes/ab-backend/internal/core/application"
	"github.com/naelcodes/ab-backend/internal/infrastructure/api"
	"github.com/naelcodes/ab-backend/pkg/server"
	"github.com/naelcodes/ab-backend/pkg/types"
)

func main() {

	config.LoadEnvironmentConfig()

	appEngine := new(server.AppEngine)
	appEngine.Init()

	context := context.Background()
	globalContext := new(types.GlobalContext)

	globalContext.Database = database.PostgresConnection(context)
	globalContext.AppEngine = appEngine
	globalContext.Context = context

	application := new(application.Application)
	application.Init(globalContext)

	restController := new(api.RestController)
	restController.Init(globalContext)
	restController.ApplicationService = application

	log.Fatal(appEngine.Start())
}
