package parser

import (
	"github.com/naelcodes/ab-backend/internal/ent"
	"github.com/naelcodes/ab-backend/internal/modules/customers/dto"
)

func CustomerModelToDTO(customer *ent.Customer) *dto.GetCustomerDTO {
	customerDTO := new(dto.GetCustomerDTO)

	customerDTO.Id = uint(customer.ID)
	customerDTO.Customer_name = customer.CustomerName
	customerDTO.State = customer.State
	customerDTO.Account_number = customer.AccountNumber
	customerDTO.Id_country = customer.IDCountry
	customerDTO.Alias = customer.Alias
	customerDTO.Ab_key = customer.AbKey
	customerDTO.Tmc_client_number = customer.TmcClientNumber

	return customerDTO
}

func CustomerModelListToDTOList(customers []*ent.Customer) []*dto.GetCustomerDTO {
	customerDTOList := make([]*dto.GetCustomerDTO, 0)

	for _, customer := range customers {
		customerDTOList = append(customerDTOList, CustomerModelToDTO(customer))
	}

	return customerDTOList
}
