package ports

import (
	"github.com/naelcodes/ab-backend/internal/common"
	PostgresAdapter "github.com/naelcodes/ab-backend/internal/modules/customers/adapters/postgres-adapter"
)

type ICustomerRepository interface {
	common.IRepository[PostgresAdapter.CustomerModel]
	GetAllCountries() []*PostgresAdapter.CountryModel
}
