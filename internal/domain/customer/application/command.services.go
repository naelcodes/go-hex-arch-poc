package application

import (
	"github.com/naelcodes/ab-backend/internal/common"
	"github.com/naelcodes/ab-backend/internal/modules/customers/dto"
)

func (application *CustomerApplication) CreateCustomerService(customer *dto.CreateCustomerDTO) (common.EID, error) {
	// TODO
	return 0, nil
}

func (application *CustomerApplication) UpdateCustomerService(customer *dto.UpdateCustomerDTO) (bool, error) {
	// TODO
	return false, nil
}

func (application *CustomerApplication) DeleteCustomerService(id common.EID) (bool, error) {
	// TODO
	return false, nil
}
