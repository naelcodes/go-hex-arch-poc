package common

import (
	"context"

	"github.com/naelcodes/ab-backend/internal/ent"
	"github.com/naelcodes/ab-backend/internal/pkg/server"
)

type GlobalContext struct {
	Database  *ent.Client
	Context   context.Context
	AppEngine *server.AppEngine
}
