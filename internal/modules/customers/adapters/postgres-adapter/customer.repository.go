package postgresAdapter

import (
	"context"
	"fmt"
	"slices"

	"github.com/naelcodes/ab-backend/internal/common"
	"github.com/naelcodes/ab-backend/internal/ent"
	"github.com/naelcodes/ab-backend/internal/ent/customer"
	"github.com/naelcodes/ab-backend/internal/modules/customers/domain"
	"github.com/naelcodes/ab-backend/internal/modules/customers/parser"
)

type CustomerRepository struct {
	Database *ent.Client
	Context  context.Context
}

func (repo *CustomerRepository) GetById(id uint) (*domain.CustomerAggregate, error) {
	return nil, nil
}

func (repo *CustomerRepository) GetAll(query *common.GetQueryParams) ([]*domain.CustomerAggregate, error) {

	customerQuery := repo.Database.Customer.Query()

	if query != nil {

		if query.Fields != nil && slices.Contains(*query.Fields, "id") && slices.Contains(*query.Fields, "name") {
			customerQuery.Select(customer.FieldID, customer.FieldCustomerName)
		}

		if query.PageNumber != nil && query.PageSize != nil {
			pageNumber := *query.PageNumber
			numberOfRecordsPerPage := *query.PageSize
			customerQuery.Offset(pageNumber * numberOfRecordsPerPage)
			customerQuery.Limit(numberOfRecordsPerPage)
		}

	}

	customers, err := customerQuery.All(repo.Context)

	if err != nil {
		fmt.Println("db-error", err)
		return nil, err
	}

	fmt.Println("Customers:", customers)
	customerAggregateList := parser.CustomerModelListToAggregateList(customers)

	return customerAggregateList, nil
}

func (repo *CustomerRepository) Count() (*int, error) {
	totalRowCount, err := repo.Database.Customer.Query().Count(repo.Context)

	if err != nil {
		return nil, err
	}
	return &totalRowCount, nil
}
func (repo *CustomerRepository) Save(customer *domain.CustomerAggregate) error {
	return nil
}
func (repo *CustomerRepository) Update(entity *CustomerModel) error {
	return nil
}
func (repo *CustomerRepository) Delete(entity *CustomerModel) error {
	return nil
}

func (repo *CustomerRepository) GetAllCountries() []*domain.CountryVO {
	return nil
}
