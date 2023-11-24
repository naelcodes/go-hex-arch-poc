package main

import (
	"context"
	"log"

	"github.com/naelcodes/ab-backend/internal/bootstrap"
	"github.com/naelcodes/ab-backend/internal/common"
	"github.com/naelcodes/ab-backend/internal/configs/database"
	"github.com/naelcodes/ab-backend/internal/pkg/server"
)

func main() {

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
