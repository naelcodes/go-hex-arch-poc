package application

import "github.com/naelcodes/ab-backend/internal/modules/customers/domain"

type appSvc struct {
	repository domain.ICustomerRepository
}

func (a *appSvc) Init(repo domain.ICustomerRepository) {
	a.repository = repo
}
