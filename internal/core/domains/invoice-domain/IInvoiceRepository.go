package invoiceDomain

import (
	"github.com/naelcodes/ab-backend/ent"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type IInvoiceRepository interface {
	Count() (*int, error)
	Exists(idInvoice types.EID) (bool, error)
	CountByCustomerId(customerId types.EID) (*int, error)
	GetByCustomerID(types.EID, *types.GetQueryParams, bool) (*dto.GetCustomerInvoicesDTO, error)
	GetById(id types.EID, query *types.GetQueryParams) (*dto.GetInvoiceDTO, error)
	GetAll(*types.GetQueryParams) (*dto.GetAllInvoiceDTO, error)
	Save(*ent.Tx, *Invoice) (*dto.GetInvoiceDTO, error)
	SaveImputation(*ent.Tx, *Invoice) error
	Update(*ent.Tx, *Invoice) error
	Void(id types.EID) error
}
