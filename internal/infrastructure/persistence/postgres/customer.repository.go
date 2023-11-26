package postgres

import (
	"context"
	"slices"

	"github.com/naelcodes/ab-backend/ent"
	"github.com/naelcodes/ab-backend/ent/customer"
	"github.com/naelcodes/ab-backend/internal/core/domain"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type CustomerRepository struct {
	Database *ent.Client
	Context  context.Context
}

func (repo *CustomerRepository) GetById(id types.EID) (*dto.GetCustomerDTO, error) {
	customer, err := repo.Database.Customer.Query().Where(customer.IDEQ(int(id))).First(repo.Context)
	if err != nil {
		return nil, err
	}
	customerDTO := CustomerModelToDTO(customer)
	return customerDTO, nil
}

func (repo *CustomerRepository) GetAll(query *types.GetQueryParams) ([]*dto.GetCustomerDTO, error) {

	customerQuery := repo.Database.Customer.Query().Where(customer.TagEQ(customer.Tag3))

	if query != nil {

		if query.Fields != nil && slices.Contains(*query.Fields, "id") && slices.Contains(*query.Fields, "name") {
			customerQuery.Select(customer.FieldID, customer.FieldCustomerName)
		}

		if query.PageNumber != nil && query.PageSize != nil {
			customerQuery.Offset(*query.PageNumber * *query.PageSize).Limit(*query.PageSize)
		}

	}

	customers, err := customerQuery.All(repo.Context)

	if err != nil {
		return nil, err
	}

	customerDTOList := CustomerModelListToDTOList(customers)

	return customerDTOList, nil
}

func (repo *CustomerRepository) Count() (*int, error) {
	totalRowCount, err := repo.Database.Customer.Query().Count(repo.Context)

	if err != nil {
		return nil, err
	}
	return &totalRowCount, nil
}

func (repo *CustomerRepository) GetCustomerOpenPayments(id types.EID) (*dto.GetCustomerOpenPaymentsDTO, error) {
	// TODO
	return nil, nil
}

func (repo *CustomerRepository) GetCustomerUnPaidInvoices(id types.EID) (*dto.GetCustomerUnpaidInvoicesDTO, error) {
	// TODO
	return nil, nil
}

func (repo *CustomerRepository) Save(customer *domain.CustomerAggregate) error {

	return nil
}
func (repo *CustomerRepository) Update(customer *domain.CustomerAggregate) error {
	return nil
}
func (repo *CustomerRepository) Delete(id types.EID) error {
	return nil
}
