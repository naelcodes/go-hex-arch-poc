package application

import (
	"github.com/naelcodes/ab-backend/internal/common"
	"github.com/naelcodes/ab-backend/internal/modules/customers/dto"
)

func (application *CustomerApplication) GetCustomerService(id common.EID) (*dto.GetCustomerDTO, error) {

	customer, err := application.ReadRepository.GetById(common.EID(id))
	if err != nil {
		return nil, err
	}
	return customer, nil
}

func (application *CustomerApplication) GetAllCustomersService(queryParams *common.GetQueryParams) (*dto.GetAllCustomersDTO, error) {

	customers, err := application.ReadRepository.GetAll(queryParams)

	if err != nil {
		return nil, err
	}

	totalRowCount, err := application.ReadRepository.Count()

	if err != nil {
		return nil, err
	}

	if queryParams == nil || queryParams.PageNumber == nil && queryParams.PageSize == nil {
		if queryParams == nil {
			queryParams = new(common.GetQueryParams)
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

func (application *CustomerApplication) GetCustomerOpenPaymentsService(id common.EID) (*dto.GetCustomerOpenPaymentsDTO, error) {
	// TODO
	return nil, nil
}

func (application *CustomerApplication) GetCustomerUnPaidInvoicesService(id common.EID) (*dto.GetCustomerUnpaidInvoicesDTO, error) {
	// TODO
	return nil, nil
}
