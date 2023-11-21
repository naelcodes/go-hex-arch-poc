package ports

import (
	"github.com/naelcodes/ab-backend/internal/common"
	PostgresAdapter "github.com/naelcodes/ab-backend/internal/modules/payments/adapters/postgres-adapter"
)

type IPaymentRepository interface {
	common.IRepository[PostgresAdapter.PaymentModel]
}
