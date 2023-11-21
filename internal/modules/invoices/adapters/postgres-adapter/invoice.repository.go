package postgresAdapter

import (
	"github.com/naelcodes/ab-backend/internal/common"
)

type InvoiceRepository struct {
	Database any
}

func (iRepo InvoiceRepository) GetById(id uint) (*InvoiceModel, error) {
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

	return nil
}
func (iRepo InvoiceRepository) Update(entity *InvoiceModel) error {
	return nil
}
func (iRepo InvoiceRepository) Delete(entity *InvoiceModel) error {
	return nil
}

func (iRepo InvoiceRepository) GetAllTravelItems() []*TravelItemModel {

	return nil
}
