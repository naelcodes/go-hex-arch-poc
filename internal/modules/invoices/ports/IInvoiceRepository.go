package ports

import (
	"github.com/naelcodes/ab-backend/internal/common"
	PostgresAdapter "github.com/naelcodes/ab-backend/internal/modules/invoices/adapters/postgres-adapter"
)

type IInvoiceRepository interface {
	common.IRepository[PostgresAdapter.InvoiceModel]
	GetAllTravelItems() []*PostgresAdapter.TravelItemModel
}
