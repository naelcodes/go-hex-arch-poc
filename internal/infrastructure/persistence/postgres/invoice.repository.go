package postgres

import (
	"context"

	"github.com/naelcodes/ab-backend/ent"
)

type InvoiceRepository struct {
	Database *ent.Client
	Context  context.Context
}
