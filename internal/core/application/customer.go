package application

import (
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
		*queryParams.PageSize = 0
	}

	getCustomersDTO := &dto.GetAllCustomersDTO{
		Data:          customers,
		PageNumber:    *queryParams.PageNumber,
		PageSize:      *queryParams.PageSize,
		TotalRowCount: *totalRowCount,
	}

	return getCustomersDTO, nil

}

func (application *Application) GetCustomerOpenPaymentsService(id types.EID) (*dto.GetCustomerOpenPaymentsDTO, error) {
	// TODO
	return nil, nil
}

func (application *Application) GetCustomerUnPaidInvoicesService(id types.EID) (*dto.GetCustomerUnpaidInvoicesDTO, error) {
	// TODO
	return nil, nil
}

func (application *Application) CreateCustomerService(customer *dto.CreateCustomerDTO) (types.EID, error) {
	// TODO
	return 0, nil
}

func (application *Application) UpdateCustomerService(customer *dto.UpdateCustomerDTO) (bool, error) {
	// TODO
	return false, nil
}

func (application *Application) DeleteCustomerService(id types.EID) (bool, error) {
	// TODO
	return false, nil
}
