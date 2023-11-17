package application

import "github.com/naelcodes/ab-backend/internal/modules/customers/dto"

type CreateCustomerService struct {
	appSvc
}

func (appSvc *CreateCustomerService) Execute(dto.CreateUserDTO) (int, error) {
	result := appSvc.repository.Save(dto)
	return result, nil
}
