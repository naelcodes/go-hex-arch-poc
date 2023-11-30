package application

import (
	builder "github.com/naelcodes/ab-backend/internal/core/application/builders"
	"github.com/naelcodes/ab-backend/internal/core/domains"
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

func (application *Application) UpdateCustomerService(customerDTO *dto.UpdateCustomerDTO) (bool, error) {

	err := application.customerRepository.Update(customerDTO)

	if err != nil {
		return false, err
	}
	return true, nil
}

func (application *Application) DeleteCustomerService(id types.EID) (bool, error) {
	domainService := new(domains.DomainService)
	err := domainService.RemoveCustomer(id, application.customerRepository, application.invoiceRepository, application.paymentRepository)

	if err != nil {
		return false, err
	}

	return true, nil
}
