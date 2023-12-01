package application

import (
	builder "github.com/naelcodes/ab-backend/internal/core/application/builders"
	customerDomain "github.com/naelcodes/ab-backend/internal/core/domains/customer-domain"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	"github.com/naelcodes/ab-backend/pkg/types"
)

func (application *Application) GetCustomerService(id types.EID) (*dto.GetCustomerDTO, error) {

	customer, err := application.customerRepository.GetById(types.EID(id))
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (application *Application) GetAllCustomersService(queryParams *types.GetQueryParams) (*dto.GetAllCustomersDTO, error) {

	customers, err := application.customerRepository.GetAll(queryParams)

	if err != nil {
		return nil, err
	}

	totalRowCount, err := application.customerRepository.Count()

	if err != nil {
		return nil, err
	}

	if queryParams == nil || (queryParams.PageNumber == nil && queryParams.PageSize == nil) {
		if queryParams == nil {
			queryParams = new(types.GetQueryParams)
		}
		queryParams.PageNumber = new(int)
		queryParams.PageSize = new(int)
		*queryParams.PageNumber = 0
		*queryParams.PageSize = *totalRowCount
	}

	getCustomersDTO := &dto.GetAllCustomersDTO{
		Data:          customers,
		PageNumber:    *queryParams.PageNumber,
		PageSize:      *queryParams.PageSize,
		TotalRowCount: *totalRowCount,
	}

	return getCustomersDTO, nil

}

func (application *Application) CreateCustomerService(customerDTO *dto.CreateCustomerDTO) (*dto.GetCustomerDTO, error) {

	customerAggregateBuilder := builder.NewCustomerAggregateBuilder()
	customerAggregateBuilder.SetCustomerName(customerDTO.CustomerName)
	customerAggregateBuilder.SetAlias(customerDTO.Alias)
	customerAggregateBuilder.SetAbKey()
	customerAggregateBuilder.SetTmcClientNumber(customerDTO.TmcClientNumber)
	customerAggregateBuilder.SetAccountNumber(customerDTO.AccountNumber)
	customerAggregateBuilder.SetState(customerDTO.State)
	customerAggregate := customerAggregateBuilder.Build()

	newCustomerDTO, err := application.customerRepository.Save(customerAggregate)

	if err != nil {
		return nil, err
	}
	return newCustomerDTO, nil
}

func (application *Application) UpdateCustomerService(id int, customerDTO *dto.UpdateCustomerDTO) (bool, error) {

	customerAggregateBuilder := builder.NewCustomerAggregateBuilder()

	customerAggregateBuilder.SetId(types.EID(id))
	if customerDTO.Customer_name != nil {
		customerAggregateBuilder.SetCustomerName(*customerDTO.Customer_name)
	}

	if customerDTO.Alias != nil {
		customerAggregateBuilder.SetAlias(*customerDTO.Alias)
	}

	if customerDTO.Tmc_client_number != nil {
		customerAggregateBuilder.SetTmcClientNumber(*customerDTO.Tmc_client_number)
	}

	if customerDTO.Account_number != nil {
		customerAggregateBuilder.SetAccountNumber(*customerDTO.Account_number)
	}

	if customerDTO.State != nil {
		customerAggregateBuilder.SetState(*customerDTO.State)
	}

	customerAggregate := customerAggregateBuilder.Build()
	err := application.customerRepository.Update(customerAggregate)

	if err != nil {
		return false, err
	}
	return true, nil
}

func (application *Application) DeleteCustomerService(id types.EID) (bool, error) {
	CustomerDomainService := &customerDomain.CustomerDomainService{
		CustomerRepository: application.customerRepository,
		InvoiceRepository:  application.invoiceRepository,
		PaymentRepository:  application.paymentRepository,
	}
	err := CustomerDomainService.RemoveCustomer(id)

	if err != nil {
		return false, err
	}

	return true, nil
}
