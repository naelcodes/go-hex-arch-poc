package ports

import (
	"github.com/naelcodes/ab-backend/internal/core/dto"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type IApplication interface {
	CreateCustomerService(dto.CreateCustomerDTO) (types.EID, error)
	UpdateCustomerService(dto.UpdateCustomerDTO) (bool, error)
	DeleteCustomerService(types.EID) (bool, error)
	GetAllCustomersService(*types.GetQueryParams) (*dto.GetAllCustomersDTO, error)
	GetCustomerService(types.EID) (*dto.GetCustomerDTO, error)
	GetCustomerPaymentsService(id types.EID) (*dto.GetCustomerPaymentsDTO, error)
	GetCustomerInvoicesService(id types.EID) (*dto.GetCustomerInvoicesDTO, error)
}
