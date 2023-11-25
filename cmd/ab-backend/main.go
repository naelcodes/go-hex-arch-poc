package main

import (
	"context"
	"log"

	"github.com/naelcodes/ab-backend/config/database"
	"github.com/naelcodes/ab-backend/ent"
	"github.com/naelcodes/ab-backend/internal/pkg/server"
	"github.com/naelcodes/ab-backend/pkg/common"
)

type GlobalContext struct {
	Database  *ent.Client
	Context   context.Context
	AppEngine *server.AppEngine
}

func main() {

	configs.LoadEnvironmentConfig()
	context := context.Background()

	appEngine := new(server.AppEngine)
	appEngine.Init()

	globalContext := new(common.GlobalContext)

	globalContext.Database = database.PostgresConnection(context, appEngine.GetLogger())
	globalContext.AppEngine = appEngine
	globalContext.Context = context

	bootstrap.InitModules(globalContext)

	log.Fatal(appEngine.Serve())
}
