package common

import "github.com/naelcodes/ab-backend/internal/pkg/server"

type GlobalContext struct {
	Database  any
	AppEngine *server.AppEngine
}

// type PersistenceContext struct {
// 	Database any
// }

type ModuleContext struct {
	Repository IRepository[any]
}
