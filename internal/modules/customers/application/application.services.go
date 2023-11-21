package application

import (
	"fmt"

	"github.com/naelcodes/ab-backend/internal/common"
	"github.com/naelcodes/ab-backend/internal/modules/customers/dto"
	"github.com/naelcodes/ab-backend/internal/modules/customers/parser"
)

func (application *CustomerApplication) GetCustomerService() {

}

func (application *CustomerApplication) GetAllCustomersService(queryParms *common.GetQueryParams) (*dto.GetCustomersDTO, error) {

	customerAggregateList, getAllError := application.repository.GetAll(queryParms)

	if getAllError != nil {
		fmt.Println("get-all-error", getAllError)
		return nil, getAllError
	}

	totalRowCount, countError := application.repository.Count()

	if countError != nil {
		return nil, countError
	}

	customersDTO := parser.CustomerAggregateListToDTOList(customerAggregateList)

	getCustomersDTO := new(dto.GetCustomersDTO)
	getCustomersDTO.Data = customersDTO
	getCustomersDTO.PageNumber = *queryParms.PageNumber
	getCustomersDTO.PageSize = *queryParms.PageSize
	getCustomersDTO.TotalRowCount = *totalRowCount

	return getCustomersDTO, nil

}

func (application *CustomerApplication) CreateCustomerService() {

}

func (application *CustomerApplication) UpdateCustomerService() {

}

func (application *CustomerApplication) DeleteCustomerService() {

}

func (application *CustomerApplication) GetAllCountriesService() {

}
