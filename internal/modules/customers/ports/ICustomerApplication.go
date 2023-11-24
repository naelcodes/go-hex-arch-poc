package ports

import (
	"github.com/naelcodes/ab-backend/internal/common"
	"github.com/naelcodes/ab-backend/internal/modules/customers/dto"
)

type ICustomerCommandService interface {
	CreateCustomerService(dto.CreateCustomerDTO) (common.EID, error)
	UpdateCustomerService(dto.UpdateCustomerDTO) (bool, error)
	DeleteCustomerService(common.EID) (bool, error)
}

type ICustomerQueryService interface {
	GetAllCustomersService(*common.GetQueryParams) (*dto.GetAllCustomersDTO, error)
	GetCustomerService(common.EID) (*dto.GetCustomerDTO, error)
	GetCustomerOpenPaymentsService(id common.EID) (*dto.GetCustomerOpenPaymentsDTO, error)
	GetCustomerUnPaidInvoicesService(id common.EID) (*dto.GetCustomerUnpaidInvoicesDTO, error)
}

type ICustomerApplication interface {
	ICustomerCommandService
	ICustomerQueryService
}
