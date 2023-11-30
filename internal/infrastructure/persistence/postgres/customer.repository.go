package postgres

import (
	"context"
	"errors"
	"fmt"
	"slices"

	"github.com/naelcodes/ab-backend/ent"
	"github.com/naelcodes/ab-backend/ent/customer"
	customerDomain "github.com/naelcodes/ab-backend/internal/core/domains/customer-domain"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	CustomErrors "github.com/naelcodes/ab-backend/pkg/errors"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type CustomerRepository struct {
	Database *ent.Client
	Context  context.Context
}

func (repo *CustomerRepository) GetById(id types.EID) (*dto.GetCustomerDTO, error) {
	customer, err := repo.Database.Customer.Query().Where(customer.IDEQ(int(id))).First(repo.Context)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, CustomErrors.RepositoryError(errors.New("customer record not found"))
		}
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error getting customer record: %v", err))
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
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error getting customers records: %v", err))
	}

	customerDTOList := CustomerModelListToDTOList(customers)

	return customerDTOList, nil
}

func (repo *CustomerRepository) Count() (*int, error) {
	totalRowCount, err := repo.Database.Customer.Query().Where(customer.TagEQ(customer.Tag3)).Count(repo.Context)

	if err != nil {
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error counting customers: %v", err))
	}
	return &totalRowCount, nil
}

func (repo *CustomerRepository) Save(customer *customerDomain.CustomerAggregate) (*dto.GetCustomerDTO, error) {

	customerRecord, err := repo.Database.Customer.Create().
		SetCustomerName(customer.CustomerName).
		SetAbKey(customer.AbKey).
		SetAlias(customer.Alias).
		SetTmcClientNumber(customer.TmcClientNumber).
		SetAccountNumber(customer.AccountNumber).
		SetState(customer.State).
		Save(repo.Context)
	if err != nil {
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error saving customer record: %v", err))
	}

	customerDTO := CustomerModelToDTO(customerRecord)
	return customerDTO, nil
}
func (repo *CustomerRepository) Update(customer *dto.UpdateCustomerDTO) error {

	UpdateQueryBuilder := repo.Database.Customer.UpdateOneID(int(customer.Id))

	if customer.Customer_name != nil {
		UpdateQueryBuilder.SetCustomerName(*customer.Customer_name)
	}

	if customer.Alias != nil {
		UpdateQueryBuilder.SetAlias(*customer.Alias)
	}

	if customer.Tmc_client_number != nil {
		UpdateQueryBuilder.SetTmcClientNumber(*customer.Tmc_client_number)
	}

	if customer.Account_number != nil {
		UpdateQueryBuilder.SetAccountNumber(*customer.Account_number)
	}

	if customer.State != nil {
		UpdateQueryBuilder.SetState(*customer.State)
	}

	_, err := UpdateQueryBuilder.Save(repo.Context)

	if err != nil {
		return CustomErrors.RepositoryError(fmt.Errorf("error updating customer record: %v", err))
	}

	return nil
}
func (repo *CustomerRepository) Delete(id types.EID) error {

	deletedCount, err := repo.Database.Customer.Delete().Where(customer.IDEQ(int(id))).Exec(repo.Context)

	if err != nil {
		return CustomErrors.RepositoryError(fmt.Errorf("error deleting customer record: %v", err))
	}

	if deletedCount == 0 {
		return CustomErrors.RepositoryError(errors.New("customer record not found"))
	}
	return nil
}
