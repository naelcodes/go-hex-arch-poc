package ports

import (
	"github.com/naelcodes/ab-backend/internal/common"
	"github.com/naelcodes/ab-backend/internal/modules/customers/dto"
)

type ICustomerReadRepository interface {
	Count() (*int, error)
	GetAll(*common.GetQueryParams) ([]*dto.GetCustomerDTO, error)
	GetById(common.EID) (*dto.GetCustomerDTO, error)
	GetCustomerOpenPayments(id common.EID) (*dto.GetCustomerOpenPaymentsDTO, error)
	GetCustomerUnPaidInvoices(id common.EID) (*dto.GetCustomerUnpaidInvoicesDTO, error)
}
