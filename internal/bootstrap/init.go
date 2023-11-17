package bootstrap

import (
	"github.com/naelcodes/ab-backend/internal/modules/customers/adapters/postgres"
	"github.com/naelcodes/ab-backend/internal/modules/customers/adapters/rest"
	"github.com/naelcodes/ab-backend/internal/pkg/database"
	"github.com/naelcodes/ab-backend/internal/pkg/server"
)

// Inject module dependencies
func InitModules(engine *server.Engine) {

	DB := database.PostgresConnection()

	customerModule := new(rest.CostumerModule)

	//pass the db here
	customerRepository := postgres.CustomerRepository{DB: DB}
	customerModule.Init(engine, customerRepository)

}
