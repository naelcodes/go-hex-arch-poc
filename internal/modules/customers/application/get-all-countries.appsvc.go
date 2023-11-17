package application

import (
	"github.com/naelcodes/ab-backend/internal/modules/customers/adapters/postgres"
	"github.com/naelcodes/ab-backend/internal/modules/customers/domain"
)

type GetAllCountriesAppSvc struct {
	repository domain.ICustomerRepository
}

func (appSvc *GetAllCountriesAppSvc) Init(repo domain.ICustomerRepository) {
	appSvc.repository = repo
}

func (appSvc *GetAllCountriesAppSvc) Execute() ([]postgres.CountryModel, error) {
	result := appSvc.repository.GetAllCountries()
	return result, nil
}
