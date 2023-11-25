package ports

import (
	InputDto "github.com/naelcodes/ab-backend/internal/dto/input"
	OutputDto "github.com/naelcodes/ab-backend/internal/dto/output"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type IApplication interface {
	CreateCustomerService(InputDto.CreateCustomerDTO) (types.EID, error)
	UpdateCustomerService(InputDto.UpdateCustomerDTO) (bool, error)
	DeleteCustomerService(types.EID) (bool, error)
	GetAllCustomersService(*types.GetQueryParams) (*OutputDto.GetAllCustomersDTO, error)
	GetCustomerService(types.EID) (*OutputDto.GetCustomerDTO, error)
	GetCustomerOpenPaymentsService(id types.EID) (*OutputDto.GetCustomerOpenPaymentsDTO, error)
	GetCustomerUnPaidInvoicesService(id types.EID) (*OutputDto.GetCustomerUnpaidInvoicesDTO, error)
}
