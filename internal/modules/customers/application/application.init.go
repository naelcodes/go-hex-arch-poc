package application

import (
	"github.com/naelcodes/ab-backend/internal/common"
	postgresAdapter "github.com/naelcodes/ab-backend/internal/modules/customers/adapters/postgres-adapter"
	"github.com/naelcodes/ab-backend/internal/modules/customers/ports"
)

type CustomerApplication struct {
	WriteRepository ports.ICustomerWriteRepository
	ReadRepository  ports.ICustomerReadRepository
	ports.ICustomerCommandService
	ports.ICustomerQueryService
}

func (application *CustomerApplication) Init(globalContext *common.GlobalContext) {

	customerReadRepository := &postgresAdapter.CustomerReadRepository{Context: globalContext.Context, Database: globalContext.Database}
	customerWriteRepository := &postgresAdapter.CustomerWriteRepository{Context: globalContext.Context, Database: globalContext.Database}

	application.WriteRepository = customerWriteRepository
	application.ReadRepository = customerReadRepository

	customerCommandApplication := new(CustomerCommandApplication)
	customerCommandApplication.CommandRepository = customerWriteRepository
	customerQueryApplication := new(CustomerQueryApplication)
	customerQueryApplication.QueryRepository = customerReadRepository

	application.Command = customerCommandApplication
	application.Query = customerQueryApplication

}
