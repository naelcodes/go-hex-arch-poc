package postgres

import (
	"github.com/naelcodes/ab-backend/ent"
	"github.com/naelcodes/ab-backend/internal/core/dto"
)

func CustomerModelToDTO(customer *ent.Customer) *dto.GetCustomerDTO {
	customerDTO := new(dto.GetCustomerDTO)

	customerDTO.Id = customer.ID
	customerDTO.Customer_name = customer.CustomerName
	customerDTO.State = customer.State
	customerDTO.Account_number = customer.AccountNumber
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
