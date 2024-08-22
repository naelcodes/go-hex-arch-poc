package postgres

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/naelcodes/ab-backend/ent"
	"github.com/naelcodes/ab-backend/ent/customer"
	customerDomain "github.com/naelcodes/ab-backend/internal/core/domains/customer-domain"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	CustomErrors "github.com/naelcodes/ab-backend/pkg/errors"
	"github.com/naelcodes/ab-backend/pkg/types"
	"github.com/naelcodes/ab-backend/pkg/utils"
)

type CustomerRepository struct {
	Database *ent.Client
	Context  context.Context
}

func (repo *CustomerRepository) GetById(id types.EID) (*dto.GetCustomerDTO, error) {
	customer, err := repo.Database.Customer.Query().Where(customer.IDEQ(int(id))).First(repo.Context)
	if err != nil {

		utils.Logger.Error(fmt.Sprintf("[CustomerRepository - GetById] Error getting customer record: %v", err))

		if ent.IsNotFound(err) {
			return nil, CustomErrors.RepositoryError(errors.New("customer record not found"))
		}
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error getting customer record: %v", err))
	}

	utils.Logger.Info(fmt.Sprintf("[CustomerRepository - GetById] Customer record: %v", customer))

	customerDTO := CustomerModelToDTO(customer)

	utils.Logger.Info(fmt.Sprintf("[CustomerRepository - GetById] Customer DTO: %v", customerDTO))

	return customerDTO, nil
}

func (repo *CustomerRepository) GetAll(queryParams *types.GetQueryParams) ([]*dto.GetCustomerDTO, error) {

	customerQuery := repo.Database.Customer.Query().Where(customer.TagEQ(customer.Tag3))

	utils.Logger.Info(fmt.Sprintf("[CustomerRepository - GetAll] QueryParams: %v", queryParams))

	if queryParams != nil {

		utils.Logger.Info(fmt.Sprintf("[CustomerRepository - GetAll] Search: %v", queryParams.Fields))

		if queryParams.Fields != nil {
			fields := *(*queryParams).Fields
			utils.Logger.Info(fmt.Sprintf("[CustomerRepository - GetAll] Fields: %v", fields))

			if strings.Contains(fields[0], "id") {
				customerQuery.Select(customer.FieldID)
			}
			if strings.Contains(fields[0], "name") {
				customerQuery.Select(customer.FieldCustomerName)
			}

		}

		utils.Logger.Info(fmt.Sprintf("[CustomerRepository - GetAll] PageNumber: %v", *queryParams.PageNumber))
		utils.Logger.Info(fmt.Sprintf("[CustomerRepository - GetAll] PageSize: %v", *queryParams.PageSize))

		if queryParams.PageNumber != nil && queryParams.PageSize != nil {
			customerQuery.Offset(*queryParams.PageNumber * *queryParams.PageSize).Limit(*queryParams.PageSize)
		}

	}

	customers, err := customerQuery.All(repo.Context)

	if err != nil {

		utils.Logger.Error(fmt.Sprintf("[CustomerRepository - GetAll] Error getting customers records: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error getting customers records: %v", err))
	}

	utils.Logger.Info(fmt.Sprintf("[CustomerRepository - GetAll] Customers records: %v", customers))
	customerDTOList := CustomerModelListToDTOList(customers)
	utils.Logger.Info(fmt.Sprintf("[CustomerRepository - GetAll] Customers DTO: %v", customerDTOList))

	return customerDTOList, nil
}

func (repo *CustomerRepository) Count() (*int, error) {

	utils.Logger.Info("[CustomerRepository - Count] counting customers")

	totalRowCount, err := repo.Database.Customer.Query().Where(customer.TagEQ(customer.Tag3)).Count(repo.Context)

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[CustomerRepository - Count] Error counting customers: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error counting customers: %v", err))
	}

	utils.Logger.Info(fmt.Sprintf("[CustomerRepository - Count] Total row count: %v", totalRowCount))
	return &totalRowCount, nil
}

func (repo *CustomerRepository) Save(customer *customerDomain.Customer) (*dto.GetCustomerDTO, error) {

	customerRecord, err := repo.Database.Customer.Create().
		SetCustomerName(customer.CustomerName).
		SetAbKey(customer.AbKey).
		SetAlias(customer.Alias).
		SetTmcClientNumber(customer.TmcClientNumber).
		SetAccountNumber(customer.AccountNumber).
		SetState(customer.State).
		Save(repo.Context)

	if err != nil {

		utils.Logger.Error(fmt.Sprintf("[CustomerRepository - Save] Error saving customer record: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error saving customer record: %v", err))
	}

	utils.Logger.Info(fmt.Sprintf("[CustomerRepository - Save] Customer record: %v", customerRecord))
	customerDTO := CustomerModelToDTO(customerRecord)
	utils.Logger.Info(fmt.Sprintf("[CustomerRepository - Save] Customer DTO: %v", customerDTO))

	return customerDTO, nil
}
func (repo *CustomerRepository) Update(customer *customerDomain.Customer) error {

	utils.Logger.Info(fmt.Sprintf("[CustomerRepository - Update] Customer Entity: %v", customer))

	UpdateQueryBuilder := repo.Database.Customer.UpdateOneID(int(customer.Id))

	if customer.CustomerName != "" {
		UpdateQueryBuilder.SetCustomerName(customer.CustomerName)
	}
	if customer.AbKey != "" {
		UpdateQueryBuilder.SetAbKey(customer.AbKey)
	}
	if customer.Alias != "" {
		UpdateQueryBuilder.SetAlias(customer.Alias)
	}
	if customer.TmcClientNumber != "" {
		UpdateQueryBuilder.SetTmcClientNumber(customer.TmcClientNumber)
	}
	if customer.AccountNumber != "" {
		UpdateQueryBuilder.SetAccountNumber(customer.AccountNumber)
	}
	if customer.State != "" {
		UpdateQueryBuilder.SetState(customer.State)
	}

	updatedCustomer, err := UpdateQueryBuilder.Save(repo.Context)

	if err != nil {

		utils.Logger.Error(fmt.Sprintf("[CustomerRepository - Update] Error updating customer record: %v", err))
		return CustomErrors.RepositoryError(fmt.Errorf("error updating customer record: %v", err))
	}

	utils.Logger.Info(fmt.Sprintf("[CustomerRepository - Update] Updated customer record: %v", updatedCustomer))

	return nil
}
func (repo *CustomerRepository) Delete(id types.EID) error {

	utils.Logger.Info(fmt.Sprintf("[CustomerRepository - Delete] Id: %v", id))

	deletedCount, err := repo.Database.Customer.Delete().Where(customer.IDEQ(int(id))).Exec(repo.Context)

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[CustomerRepository - Delete] Error deleting customer record: %v", err))
		return CustomErrors.RepositoryError(fmt.Errorf("error deleting customer record: %v", err))
	}

	utils.Logger.Info(fmt.Sprintf("[CustomerRepository - Delete] Deleted customer record count: %v", deletedCount))

	if deletedCount == 0 {

		utils.Logger.Error(fmt.Sprintf("[CustomerRepository - Delete] Customer record  with id %v not found", id))
		return CustomErrors.RepositoryError(errors.New("customer record not found"))
	}

	return nil
}
