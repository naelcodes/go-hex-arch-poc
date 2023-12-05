package customerDomain

import (
	"github.com/naelcodes/ab-backend/pkg/types"
	"github.com/naelcodes/ab-backend/pkg/utils"
)

type CustomerBuilder struct {
	customer *Customer
}

func NewCustomerBuilder() *CustomerBuilder {
	builder := new(CustomerBuilder)
	builder.customer = new(Customer)
	return builder
}

func (builder *CustomerBuilder) SetId(id types.EID) *CustomerBuilder {
	builder.customer.Id = id
	return builder
}

func (builder *CustomerBuilder) SetCustomerName(name string) *CustomerBuilder {
	builder.customer.CustomerName = name
	return builder
}

func (builder *CustomerBuilder) SetAlias(alias string) *CustomerBuilder {
	builder.customer.Alias = alias
	return builder
}

func (builder *CustomerBuilder) SetAbKey() *CustomerBuilder {
	builder.customer.AbKey = utils.GenerateRandomString(15)
	return builder
}

func (builder *CustomerBuilder) SetTmcClientNumber(tmcClientNumber string) *CustomerBuilder {
	builder.customer.TmcClientNumber = tmcClientNumber
	return builder
}

func (builder *CustomerBuilder) SetAccountNumber(accountNumber string) *CustomerBuilder {
	builder.customer.AccountNumber = accountNumber
	return builder
}

func (builder *CustomerBuilder) SetState(state string) *CustomerBuilder {
	builder.customer.State = state
	return builder
}

func (builder *CustomerBuilder) Build() *Customer {
	return builder.customer
}
