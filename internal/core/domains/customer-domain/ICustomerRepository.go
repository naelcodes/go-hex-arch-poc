package customerDomain

import (
	"github.com/naelcodes/ab-backend/internal/core/dto"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type ICustomerRepository interface {
	Count() (*int, error)
	GetAll(*types.GetQueryParams) ([]*dto.GetCustomerDTO, error)
	GetById(types.EID) (*dto.GetCustomerDTO, error)
	Save(*Customer) (*dto.GetCustomerDTO, error)
	Update(*Customer) error
	Delete(types.EID) error
}
