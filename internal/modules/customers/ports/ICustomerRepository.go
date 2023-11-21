package ports

import (
	"github.com/naelcodes/ab-backend/internal/common"
	"github.com/naelcodes/ab-backend/internal/modules/customers/domain"
)

type ICustomerRepository interface {
	common.IRepository[domain.CustomerAggregate]
	GetAll(*common.GetQueryParams) ([]*domain.CustomerAggregate, error)
	Count() (*int, error)
	GetAllCountries() []*domain.CountryVO
}
