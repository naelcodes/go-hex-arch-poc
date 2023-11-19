package domain

import (
	"github.com/naelcodes/ab-backend/internal/common"
	"github.com/naelcodes/ab-backend/internal/modules/invoices/adapters/postgres"
)

type IInvoiceRepository interface {
	common.IRepository[postgres.InvoiceModel]
	GetAllTravelItems() []*postgres.TravelItemModel
}
