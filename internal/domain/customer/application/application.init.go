package application

import (
	"github.com/naelcodes/ab-backend/internal/common"
	postgresAdapter "github.com/naelcodes/ab-backend/internal/modules/customers/adapters/postgres-adapter"
	"github.com/naelcodes/ab-backend/internal/modules/customers/ports"
)

type CustomerApplication struct {
	WriteRepository ports.ICustomerWriteRepository
	ReadRepository  ports.ICustomerReadRepository
}

func (application *CustomerApplication) Init(globalContext *common.GlobalContext) {

	customerReadRepository := &p.CustomerReadRepository{Context: globalContext.Context, Database: globalContext.Database}
	customerWriteRepository := &postgresAdapter.CustomerWriteRepository{Context: globalContext.Context, Database: globalContext.Database}

	application.WriteRepository = customerWriteRepository
	application.ReadRepository = customerReadRepository

	customerApplication := new(CustomerApplication)

	customerApplication.ReadRepository = customerReadRepository
	customerApplication.WriteRepository = customerWriteRepository

}
