package dto

import "github.com/naelcodes/ab-backend/internal/modules/customers/domain/entities"

type Parser struct{}

type CustomerResponseDTO interface {
	GetCustomerDTO | GetCustomersDTO | GetAllCountriesDTO
}
type convert[T CustomerResponseDTO] interface {
	convertEntityToDTO(e entities.CustomerEntity, DTOtype string) T
}

func (p Parser) convertEntityToDTO(e entities.CustomerEntity, DTOtype string) CustomerResponseDTO {
	switch DTOtype {
	case "getOne":
		return GetCustomerDTO{
			Id:            e.ID,
			Customer_name: e.Name,
			State:         e.Country.Name,
			slug:          e.Slug,
			Ab_key:        e.AbKey,
			Id_country:    uint(e.Country.Id),
		}
	}
}
