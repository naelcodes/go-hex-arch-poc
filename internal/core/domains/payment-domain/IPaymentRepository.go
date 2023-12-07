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
	SavePaymentAllocation(*ent.Tx, *Payment) error
	CheckInvoiceOwnerPayments(invoiceId types.EID, paymentIds []int) (*[]int, error)
	Delete(id types.EID) error
}
