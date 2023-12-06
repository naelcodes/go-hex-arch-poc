package invoiceDomain

import (
	"github.com/naelcodes/ab-backend/ent"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type IInvoiceRepository interface {
	Count() (*int, error)
	CountByCustomerId(customerId types.EID) (*int, error)
	GetByCustomerID(types.EID, *types.GetQueryParams, *bool) ([]*dto.GetInvoiceDTO, error)
	GetById(id types.EID, query *types.GetQueryParams) (*dto.GetInvoiceDTO, error)
	GetAll(*types.GetQueryParams) (*dto.GetAllInvoiceDTO, error)
	Save(*ent.Tx, *Invoice) (*dto.GetInvoiceDTO, error)
	SaveImputation(*ent.Tx, *Invoice)
	Update(*ent.Tx, *Invoice) error
	Void(id types.EID) error
}
