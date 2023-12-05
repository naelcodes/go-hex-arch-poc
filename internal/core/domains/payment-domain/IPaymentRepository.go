package paymentDomain

import (
	"github.com/naelcodes/ab-backend/ent"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type IPaymentRepository interface {
	Count() (*int, error)
	CountByCustomerID(customerId types.EID) (*int, error)
	GetByCustomerID(id types.EID, queryParams *types.GetQueryParams, isOpen *bool) ([]*dto.GetPaymentDTO, error)
	GetById(id types.EID) (*dto.GetPaymentDTO, error)
	GetAll(*types.GetQueryParams) ([]*dto.GetPaymentDTO, error)
	Save(*Payment) (*dto.GetPaymentDTO, error)
	Update(*Payment) error
	SaveAllPaymentsAllocations(*ent.Tx, []*Payment)
	Delete(id types.EID) error
}
