package postgresAdapter

import (
	"github.com/naelcodes/ab-backend/internal/common"
)

type CustomerRepository struct {
	Database any
}

func (cRepo CustomerRepository) FindById(id uint) (*CustomerModel, error) {
	return nil, nil
}

func (cRepo CustomerRepository) Find(query common.GetQueryParams, options ...any) ([]*CustomerModel, error) {
	return nil, nil
}
func (cRepo CustomerRepository) FindAll() ([]*CustomerModel, error) {
	return nil, nil
}
func (cRepo CustomerRepository) Count() int {
	return 0
}
func (cRepo CustomerRepository) Save(newCustomer *CustomerModel) error {
	// result := cRepo.DB.Create(&newCustomer)
	// if result.Error != nil {
	// 	return result.Error
	// }
	return nil
}
func (cRepo CustomerRepository) Update(entity *CustomerModel) error {
	return nil
}
func (cRepo CustomerRepository) Delete(entity *CustomerModel) error {
	return nil
}

func (cRepo CustomerRepository) GetAllCountries() []*CountryModel {

	// countries := []*CountryModel{}
	// cRepo.DB.Where("currency_code=?", "XOF").Find(&countries)

	return nil
}
