package ports

import (
	"github.com/naelcodes/ab-backend/internal/common"
	"github.com/naelcodes/ab-backend/internal/modules/customers/domain"
)

type ICustomerWriteRepository interface {
	Save(*domain.CustomerAggregate) error
	Update(*domain.CustomerAggregate) error
	Delete(common.EID) error
}
