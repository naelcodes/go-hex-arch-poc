package domain

import (
	"github.com/naelcodes/ab-backend/internal/common/base"
)

type ICustomerRepository interface {
	base.IRepository[CustomerEntity]
	GetAllCountries() []CountryVO
}
