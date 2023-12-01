package customerDomain

import (
	"github.com/naelcodes/ab-backend/internal/core/dto"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type ICustomerRepository interface {
	Count() (*int, error)
	GetAll(*types.GetQueryParams) ([]*dto.GetCustomerDTO, error)
	GetById(types.EID) (*dto.GetCustomerDTO, error)
	Save(*CustomerAggregate) (*dto.GetCustomerDTO, error)
	Update(*CustomerAggregate) error
	Delete(types.EID) error
}
