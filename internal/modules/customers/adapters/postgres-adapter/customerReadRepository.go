package postgresAdapter

import (
	"context"
	"errors"
	"slices"

	"github.com/naelcodes/ab-backend/internal/common"
	"github.com/naelcodes/ab-backend/internal/ent"
	"github.com/naelcodes/ab-backend/internal/ent/customer"
	"github.com/naelcodes/ab-backend/internal/modules/customers/dto"
	"github.com/naelcodes/ab-backend/internal/modules/customers/parser"
)

var (
	err = errors.New("Customer Not Found")
)

type CustomerReadRepository struct {
	Database *ent.Client
	Context  context.Context
}

func (repo *CustomerReadRepository) GetById(id common.EID) (*dto.GetCustomerDTO, error) {
	customer, err := repo.Database.Customer.Query().Where(customer.IDEQ(int(id))).First(repo.Context)
	if err != nil {
		return nil, err
	}
	customerDTO := parser.CustomerModelToDTO(customer)
	return customerDTO, nil
}

func (repo *CustomerReadRepository) GetAll(query *common.GetQueryParams) ([]*dto.GetCustomerDTO, error) {

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

	customerDTOList := parser.CustomerModelListToDTOList(customers)

	return customerDTOList, nil
}

func (repo *CustomerReadRepository) Count() (*int, error) {
	totalRowCount, err := repo.Database.Customer.Query().Count(repo.Context)

	if err != nil {
		return nil, err
	}
	return &totalRowCount, nil
}

func (repo *CustomerReadRepository) GetCustomerOpenPayments(id common.EID) (*dto.GetCustomerOpenPaymentsDTO, error) {
	// TODO
	return nil, nil
}

func (repo *CustomerReadRepository) GetCustomerUnPaidInvoices(id common.EID) (*dto.GetCustomerUnpaidInvoicesDTO, error) {
	// TODO
	return nil, nil
}
