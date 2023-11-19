package parser

// func ConvertEntityToGetCustomerDTO(e *domain.CustomerEntity) *GetCustomerDTO {

// 	return &GetCustomerDTO{
// 		Id:                e.Id,
// 		Customer_name:     e.Name,
// 		State:             e.Country.Name,
// 		Slug:              e.Slug,
// 		Ab_key:            e.AbKey,
// 		Id_country:        uint(e.Country.Id),
// 		Alias:             e.Alias,
// 		Account_number:    e.AccountNumber,
// 		Tmc_client_number: e.TmcClientNumber,
// 	}

// }

// func ConvertEntityListToGetCustomersDTO(el []*domain.CustomerEntity) []*GetCustomerDTO {
// 	customerDTOList := make([]*GetCustomerDTO, len(el))

// 	for _, e := range el {
// 		customerDTOList = append(customerDTOList, ConvertEntityToGetCustomerDTO(e))
// 	}
// 	return customerDTOList
// }

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
