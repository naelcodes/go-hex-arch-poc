package paymentDomain

import (
	"github.com/naelcodes/ab-backend/ent"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type IPaymentRepository interface {
	Count() (*int, error)
	CountByCustomerID(customerId types.EID) (*int, error)
	GetById(id types.EID) (*dto.GetPaymentDTO, error)
	GetAll(*types.GetQueryParams) (*dto.GetAllPaymentsDTO, error)
	Save(*Payment) (*dto.GetPaymentDTO, error)
	Update(*Payment) error
	SaveAllPaymentsAllocations(*ent.Tx, []*Payment)
	Delete(id types.EID) error
}
