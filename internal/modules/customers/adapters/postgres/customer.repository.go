package postgres

import (
	"github.com/naelcodes/ab-backend/internal/common/base"
	"github.com/naelcodes/ab-backend/internal/common/types"
	"github.com/naelcodes/ab-backend/internal/modules/customers/domain"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	DB *gorm.DB
}

func (cRepo CustomerRepository) FindById(id types.Id) (*domain.CustomerEntity, error) {
	return nil, nil
}

func (cRepo CustomerRepository) Find(query base.IQueryBuilder) ([]*domain.CustomerEntity, error) {
	return nil, nil
}
func (cRepo CustomerRepository) FindAll() ([]*domain.CustomerEntity, error) {
	return nil, nil
}
func (cRepo CustomerRepository) Count() int {
	return 0
}
func (cRepo CustomerRepository) Save(newCustomer *domain.CustomerEntity) error {
	result := cRepo.DB.Create(&newCustomer)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (cRepo CustomerRepository) Update(entity *domain.CustomerEntity) error {
	return nil
}
func (cRepo CustomerRepository) Delete(entity *domain.CustomerEntity) error {
	return nil
}

func (cRepo CustomerRepository) GetAllCountries() []CountryModel {

	countries := []CountryModel{}
	cRepo.DB.Where("currency_code=?", "XOF").Find(&countries)

	return countries
}
