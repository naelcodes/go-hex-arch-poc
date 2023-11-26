package domain

import (
	"github.com/naelcodes/ab-backend/internal/core/dto"
	"github.com/naelcodes/ab-backend/pkg/types"
	"github.com/naelcodes/ab-backend/pkg/utils"
)

type ICustomerRepository interface {
	Count() (*int, error)
	GetAll(*types.GetQueryParams) ([]*dto.GetCustomerDTO, error)
	GetById(types.EID) (*dto.GetCustomerDTO, error)
	GetCustomerOpenPayments(id types.EID) (*dto.GetCustomerOpenPaymentsDTO, error)
	GetCustomerUnPaidInvoices(id types.EID) (*dto.GetCustomerUnpaidInvoicesDTO, error)
	Save(CustomerAggregate) error
	Update(CustomerAggregate) error
	Delete(types.EID) error
}

type CustomerAggregate struct {
	types.BaseEntity
	Name            string
	Alias           string
	AbKey           string
	TmcClientNumber string
	AccountNumber   string
	State           string

	InvoiceIds []types.EID
	PaymentIds []types.EID
}

type CustomerAggregateBuilder struct {
	customerAggregate *CustomerAggregate
}

func NewCustomerAggregateBuilder() *CustomerAggregateBuilder {
	builder := new(CustomerAggregateBuilder)
	builder.customerAggregate = new(CustomerAggregate)
	return builder
}

func (builder *CustomerAggregateBuilder) SetName(name string) *CustomerAggregateBuilder {
	builder.customerAggregate.Name = name
	return builder
}

func (builder *CustomerAggregateBuilder) SetAlias(alias string) *CustomerAggregateBuilder {
	builder.customerAggregate.Alias = alias
	return builder
}

func (builder *CustomerAggregateBuilder) SetAbKey() *CustomerAggregateBuilder {
	builder.customerAggregate.AbKey = utils.GenerateRandomString(15)
	return builder
}

func (builder *CustomerAggregateBuilder) SetTmcClientNumber(tmcClientNumber string) *CustomerAggregateBuilder {
	builder.customerAggregate.TmcClientNumber = tmcClientNumber
	return builder
}

func (builder *CustomerAggregateBuilder) AddAccountNumber(accountNumber string) *CustomerAggregateBuilder {
	builder.customerAggregate.AccountNumber = accountNumber
	return builder
}

func (builder *CustomerAggregateBuilder) AddState(state string) *CustomerAggregateBuilder {
	builder.customerAggregate.State = state
	return builder
}

func (builder *CustomerAggregateBuilder) AddInvoices(invoiceIds []types.EID) *CustomerAggregateBuilder {
	return builder
}

func (builder *CustomerAggregateBuilder) AddPayments(paymentIds []string) *CustomerAggregateBuilder {
	//b.payments = payments
	return builder
}

func (builder *CustomerAggregateBuilder) Build() *CustomerAggregate {
	return builder.customerAggregate
}
