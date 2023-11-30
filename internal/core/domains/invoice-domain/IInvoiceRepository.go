package invoiceDomain

import (
	"github.com/naelcodes/ab-backend/internal/core/dto"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type IInvoiceRepository interface {
	Count() (int, error)
	CountByCustomerId(customerId types.EID) (int, error)
	GetByCustomerID(id types.EID, isPaid bool) ([]*dto.GetInvoiceDTO, error)
	GetById(id types.EID) (*dto.GetInvoiceDTO, error)
	GetAll(*types.GetQueryParams) ([]*dto.GetInvoiceDTO, error)
	Save(*InvoiceAggregate) error
	Update(*InvoiceAggregate) error
	Delete(id types.EID) error
}
