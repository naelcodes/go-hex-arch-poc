package main

import (
	"context"
	"log"

	"github.com/naelcodes/ab-backend/config"
	"github.com/naelcodes/ab-backend/config/database"
	"github.com/naelcodes/ab-backend/pkg/server"
	"github.com/naelcodes/ab-backend/pkg/types"
)

func main() {

	config.LoadEnvironmentConfig()

	appEngine := new(server.AppEngine)
	appEngine.Init()

	context := context.Background()
	globalContext := new(types.GlobalContext)

	globalContext.Database = database.PostgresConnection(context, appEngine.GetLogger())
	globalContext.AppEngine = appEngine
	globalContext.Context = context

	log.Fatal(appEngine.Start())
}
