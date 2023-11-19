package main

import (
	"log"

	"github.com/naelcodes/ab-backend/internal/bootstrap"
	"github.com/naelcodes/ab-backend/internal/common"
	"github.com/naelcodes/ab-backend/internal/pkg/database"
	"github.com/naelcodes/ab-backend/internal/pkg/server"
)

func main() {

	appEngine := new(server.AppEngine)
	appEngine.Init()

	globalContext := new(common.GlobalContext)
	globalContext.Database = database.PostgresConnection()
	globalContext.AppEngine = appEngine

	bootstrap.InitModules(globalContext)

	log.Fatal(appEngine.Serve())
}
