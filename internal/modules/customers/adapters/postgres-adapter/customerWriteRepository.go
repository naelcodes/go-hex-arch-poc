package postgresAdapter

import (
	"context"

	"github.com/naelcodes/ab-backend/internal/common"
	"github.com/naelcodes/ab-backend/internal/ent"
	"github.com/naelcodes/ab-backend/internal/modules/customers/domain"
)

type CustomerWriteRepository struct {
	Database *ent.Client
	Context  context.Context
}

func (repo *CustomerWriteRepository) Save(customer *domain.CustomerAggregate) error {

	return nil
}
func (repo *CustomerWriteRepository) Update(customer *domain.CustomerAggregate) error {
	return nil
}
func (repo *CustomerWriteRepository) Delete(id common.EID) error {
	return nil
}
