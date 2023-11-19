package postgres

import (
	"github.com/naelcodes/ab-backend/internal/common"
	"gorm.io/gorm"
)

type InvoiceRepository struct {
	DB *gorm.DB
}

func (iRepo InvoiceRepository) FindById(id uint) (*InvoiceModel, error) {
	return nil, nil
}

func (iRepo InvoiceRepository) Find(query common.GetQueryParams, options ...any) ([]*InvoiceModel, error) {
	return nil, nil
}
func (iRepo InvoiceRepository) FindAll() ([]*InvoiceModel, error) {
	return nil, nil
}
func (iRepo InvoiceRepository) Count() int {
	return 0
}
func (iRepo InvoiceRepository) Save(newInvoice *InvoiceModel) error {
	result := iRepo.DB.Create(&newInvoice)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (iRepo InvoiceRepository) Update(entity *InvoiceModel) error {
	return nil
}
func (iRepo InvoiceRepository) Delete(entity *InvoiceModel) error {
	return nil
}

func (iRepo InvoiceRepository) GetAllTravelItems() []*TravelItemModel {

	travelItemsList := []*TravelItemModel{}
	iRepo.DB.Where("currency_code=?", "XOF").Find(&travelItemsList)

	return travelItemsList
}
