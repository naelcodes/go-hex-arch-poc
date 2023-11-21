package parser

import (
	"github.com/naelcodes/ab-backend/internal/modules/customers/domain"
	"github.com/naelcodes/ab-backend/internal/modules/customers/dto"
)

func CustomerAggregateToDTO(customerAggregate *domain.CustomerAggregate) *dto.GetCustomerDTO {
	customerDTO := new(dto.GetCustomerDTO)

	customerDTO.Id = customerAggregate.Id
	customerDTO.Customer_name = customerAggregate.Name
	customerDTO.State = customerAggregate.State
	customerDTO.Ab_key = customerAggregate.AbKey
	customerDTO.Id_country = customerAggregate.CountryId
	customerDTO.Alias = customerAggregate.Alias
	customerDTO.Account_number = customerAggregate.AccountNumber
	customerDTO.Tmc_client_number = customerAggregate.TmcClientNumber
	return customerDTO

}

func CustomerAggregateListToDTOList(customerAggregateList []*domain.CustomerAggregate) []*dto.GetCustomerDTO {
	customerDTOList := make([]*dto.GetCustomerDTO, 0)

	for _, customerAggregate := range customerAggregateList {
		customerDTOList = append(customerDTOList, CustomerAggregateToDTO(customerAggregate))
	}
	return customerDTOList
}

// func ConvertCreateDtoToDataModel(dto *CreateCustomerDTO) *postgres.CustomerModel {
// 	return &postgres.CustomerModel{
// 		Customer_name:     dto.Customer_name,
// 		State:             dto.State,
// 		Id_country:        dto.Id_country,
// 		Alias:             dto.Alias,
// 		Account_number:    dto.Account_number,
// 		Tmc_client_number: dto.Tmc_client_number,
// 	}
// }

// func ConvertGetCountryVOToDTO(cvo []*domain.CountryVO) {}
