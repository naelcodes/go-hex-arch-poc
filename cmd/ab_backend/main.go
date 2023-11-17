package main

import (
	"log"

	"github.com/naelcodes/ab-backend/internal/bootstrap"
	"github.com/naelcodes/ab-backend/internal/pkg/server"
)

func main() {
	app := new(server.Engine)

	app.Init()
	bootstrap.InitModules(app)

	log.Fatal(app.Serve())
}
