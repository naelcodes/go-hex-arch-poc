package application

import (
	"fmt"

	customerDomain "github.com/naelcodes/ab-backend/internal/core/domains/customer-domain"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	"github.com/naelcodes/ab-backend/pkg/types"
	"github.com/naelcodes/ab-backend/pkg/utils"
)

func (application *Application) GetCustomerService(id types.EID) (*dto.GetCustomerDTO, error) {

	customer, err := application.customerRepository.GetById(types.EID(id))
	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[GetCustomerService] Error getting customer record: %v", err))
		return nil, err
	}

	utils.Logger.Info(fmt.Sprintf("[GetCustomerService] Customer record: %v", customer))
	return customer, nil
}

func (application *Application) GetAllCustomersService(queryParams *types.GetQueryParams) (*dto.GetAllCustomersDTO, error) {

	utils.Logger.Info(fmt.Sprintf("[GetAllCustomersService] QueryParams -: %v", *queryParams))

	totalRowCount, err := application.customerRepository.Count()

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[GetAllCustomersService] Error counting customers: %v", err))
		return nil, err
	}

	utils.Logger.Info(fmt.Sprintf("[GetAllCustomersService] Total number of customers: %v", totalRowCount))

	if queryParams == nil || (queryParams.PageNumber == nil && queryParams.PageSize == nil) {
		if queryParams == nil {
			queryParams = new(types.GetQueryParams)
		}
		queryParams.PageNumber = new(int)
		queryParams.PageSize = new(int)
		*queryParams.PageNumber = 0
		*queryParams.PageSize = *totalRowCount
	}

	customers, err := application.customerRepository.GetAll(queryParams)

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[GetAllCustomersService] Error getting customers: %v", err))
		return nil, err
	}

	getCustomersDTO := &dto.GetAllCustomersDTO{
		Data:          customers,
		PageNumber:    *queryParams.PageNumber,
		PageSize:      *queryParams.PageSize,
		TotalRowCount: *totalRowCount,
	}

	utils.Logger.Info(fmt.Sprintf("[GetAllCustomersService] GetCustomersDTO: %v", getCustomersDTO))

	return getCustomersDTO, nil

}

func (application *Application) CreateCustomerService(customerDTO *dto.CreateCustomerDTO) (*dto.GetCustomerDTO, error) {

	utils.Logger.Info(fmt.Sprintf("[CreateCustomerService] CreateCustomerDTO: %v", customerDTO))

	customer := customerDomain.NewCustomerBuilder().
		SetCustomerName(customerDTO.CustomerName).
		SetAlias(customerDTO.Alias).
		SetAbKey().
		SetTmcClientNumber(customerDTO.TmcClientNumber).
		SetAccountNumber(customerDTO.AccountNumber).
		SetState(customerDTO.State).
		Build()

	newCustomerDTO, err := application.customerRepository.Save(customer)

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[CreateCustomerService] Error saving customer: %v", err))
		return nil, err
	}

	utils.Logger.Info(fmt.Sprintf("[CreateCustomerService] NewCustomerDTO: %v", newCustomerDTO))
	return newCustomerDTO, nil
}

func (application *Application) UpdateCustomerService(id int, customerDTO *dto.UpdateCustomerDTO) (bool, error) {

	utils.Logger.Info(fmt.Sprintf("[UpdateCustomerService] UpdateCustomerDTO: %v", customerDTO))

	customerBuilder := customerDomain.NewCustomerBuilder()

	customerBuilder.SetId(types.EID(id))
	if customerDTO.Customer_name != nil {
		customerBuilder.SetCustomerName(*customerDTO.Customer_name)
	}

	if customerDTO.Alias != nil {
		customerBuilder.SetAlias(*customerDTO.Alias)
	}

	if customerDTO.Tmc_client_number != nil {
		customerBuilder.SetTmcClientNumber(*customerDTO.Tmc_client_number)
	}

	if customerDTO.Account_number != nil {
		customerBuilder.SetAccountNumber(*customerDTO.Account_number)
	}

	if customerDTO.State != nil {
		customerBuilder.SetState(*customerDTO.State)
	}

	customer := customerBuilder.Build()
	err := application.customerRepository.Update(customer)

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[UpdateCustomerService] Error updating customer: %v", err))
		return false, err
	}

	utils.Logger.Info(fmt.Sprintf("[UpdateCustomerService] customer updated: %v", true))

	return true, nil
}

func (application *Application) DeleteCustomerService(id types.EID) (bool, error) {

	utils.Logger.Info(fmt.Sprintf("[DeleteCustomerService] Id: %v", id))

	CustomerDomainService := customerDomain.NewCustomerDomainService(application.customerRepository, application.invoiceRepository, application.paymentRepository)
	err := CustomerDomainService.RemoveCustomer(id)

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[DeleteCustomerService] Error deleting customer: %v", err))
		return false, err
	}

	utils.Logger.Info(fmt.Sprintf("[DeleteCustomerService] customer deleted: %v", true))
	return true, nil
}

func (application *Application) GetCustomerPaymentsService(customerId types.EID, queryParams *types.GetQueryParams, isOpen *bool) (*dto.GetCustomerPaymentsDTO, error) {

	utils.Logger.Info(fmt.Sprintf("[GetCustomerPaymentsService] customerId: %v, queryParams: %v, isOpen: %v", customerId, queryParams, isOpen))

	totalRowCount, err := application.paymentRepository.CountByCustomerID(customerId)

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[GetCustomerPaymentsService] Error getting totalRowCount: %v", err))
		return nil, err
	}

	utils.Logger.Info(fmt.Sprintf("[GetCustomerPaymentsService] totalRowCount: %v", totalRowCount))

	if queryParams == nil || (queryParams.PageNumber == nil && queryParams.PageSize == nil) {
		if queryParams == nil {
			queryParams = new(types.GetQueryParams)
		}
		queryParams.PageNumber = new(int)
		queryParams.PageSize = new(int)
		*queryParams.PageNumber = 0
		*queryParams.PageSize = *totalRowCount
	}

	paymentDTO, err := application.paymentRepository.GetByCustomerID(customerId, queryParams, isOpen)

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[GetCustomerPaymentsService] Error getting paymentDTO: %v", err))
		return nil, err
	}

	customerPayments := new(dto.CustomerPayments)
	customerPayments.IdCustomer = int(customerId)
	customerPayments.Payments = paymentDTO

	getCustomPaymentsDTO := new(dto.GetCustomerPaymentsDTO)
	getCustomPaymentsDTO.Data = customerPayments
	getCustomPaymentsDTO.PageNumber = *queryParams.PageNumber
	getCustomPaymentsDTO.PageSize = *queryParams.PageSize
	getCustomPaymentsDTO.TotalRowCount = *totalRowCount

	utils.Logger.Info(fmt.Sprintf("[GetCustomerPaymentsService] getCustomPaymentsDTO: %v", getCustomPaymentsDTO))

	return getCustomPaymentsDTO, nil
}

func (application *Application) GetCustomerInvoicesService(customerId types.EID, queryParams *types.GetQueryParams, isPaid *bool) (*dto.GetCustomerInvoicesDTO, error) {

	utils.Logger.Info(fmt.Sprintf("[GetCustomerInvoicesService] customerId: %v, queryParams: %v, isPaid: %v", customerId, queryParams, isPaid))

	totalRowCount, err := application.invoiceRepository.CountByCustomerId(customerId)

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[GetCustomerInvoicesService] Error getting totalRowCount: %v", err))
		return nil, err
	}

	utils.Logger.Info(fmt.Sprintf("[GetCustomerInvoicesService] totalRowCount: %v", totalRowCount))

	if queryParams == nil || (queryParams.PageNumber == nil && queryParams.PageSize == nil) {
		if queryParams == nil {
			queryParams = new(types.GetQueryParams)
		}
		queryParams.PageNumber = new(int)
		queryParams.PageSize = new(int)
		*queryParams.PageNumber = 0
		*queryParams.PageSize = *totalRowCount
	}

	invoiceDTO, err := application.invoiceRepository.GetByCustomerID(customerId, queryParams, isPaid)

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[GetCustomerInvoicesService] Error getting invoiceDTO: %v", err))
		return nil, err
	}

	customerInvoices := new(dto.CustomerInvoice)
	customerInvoices.IdCustomer = int(customerId)
	customerInvoices.Invoices = invoiceDTO

	getCustomerInvoicesDTO := new(dto.GetCustomerInvoicesDTO)
	getCustomerInvoicesDTO.Data = customerInvoices
	getCustomerInvoicesDTO.PageNumber = *queryParams.PageNumber
	getCustomerInvoicesDTO.PageSize = *queryParams.PageSize
	getCustomerInvoicesDTO.TotalRowCount = *totalRowCount

	utils.Logger.Info(fmt.Sprintf("[GetCustomerInvoicesService] getCustomerInvoicesDTO: %v", *getCustomerInvoicesDTO))

	return getCustomerInvoicesDTO, nil

}
