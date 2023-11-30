package builder

import (
	customerDomain "github.com/naelcodes/ab-backend/internal/core/domains/customer-domain"
	"github.com/naelcodes/ab-backend/pkg/utils"
)

type CustomerAggregateBuilder struct {
	customerAggregate *customerDomain.CustomerAggregate
}

func NewCustomerAggregateBuilder() *CustomerAggregateBuilder {
	builder := new(CustomerAggregateBuilder)
	builder.customerAggregate = new(customerDomain.CustomerAggregate)
	return builder
}

func (builder *CustomerAggregateBuilder) SetCustomerName(name string) *CustomerAggregateBuilder {
	builder.customerAggregate.CustomerName = name
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

func (builder *CustomerAggregateBuilder) SetAccountNumber(accountNumber string) *CustomerAggregateBuilder {
	builder.customerAggregate.AccountNumber = accountNumber
	return builder
}

func (builder *CustomerAggregateBuilder) SetState(state string) *CustomerAggregateBuilder {
	builder.customerAggregate.State = state
	return builder
}

func (builder *CustomerAggregateBuilder) Build() *customerDomain.CustomerAggregate {
	return builder.customerAggregate
}
