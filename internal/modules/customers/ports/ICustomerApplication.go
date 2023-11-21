package ports

import (
	"github.com/naelcodes/ab-backend/internal/common"
	"github.com/naelcodes/ab-backend/internal/modules/customers/dto"
)

type ICustomerApplication interface {
	CreateCustomerService()
	UpdateCustomerService()
	GetAllCountriesService()
	GetAllCustomersService(*common.GetQueryParams) (*dto.GetCustomersDTO, error)
	GetCustomerService()
	DeleteCustomerService()
}
