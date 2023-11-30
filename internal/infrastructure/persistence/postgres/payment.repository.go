package postgres

import (
	"context"

	"github.com/naelcodes/ab-backend/ent"
)

type PaymentRepository struct {
	Database *ent.Client
	Context  context.Context
}
