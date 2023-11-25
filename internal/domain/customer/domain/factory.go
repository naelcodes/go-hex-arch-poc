package factory

import (
	"errors"

	"github.com/naelcodes/ab-backend/internal/modules/customers/domain"
	"github.com/naelcodes/ab-backend/internal/modules/customers/dto"
	"github.com/naelcodes/ab-backend/internal/pkg/utils"
)

type CustomerAggregateBuilder struct {
	name            string
	alias           string
	abKey           string
	tmcClientNumber string
	accountNumber   string
	state           string
	invoices        []string
	payments        []string
}

func NewCustomerAggregateBuilder() *CustomerAggregateBuilder {
	return &CustomerAggregateBuilder{}
}

func (b *CustomerAggregateBuilder) AddName(name string) *CustomerAggregateBuilder {
	b.name = name
	return b
}

func (b *CustomerAggregateBuilder) AddAlias(alias string) *CustomerAggregateBuilder {
	b.alias = alias
	return b
}

func (b *CustomerAggregateBuilder) AddAbKey(abKey string) *CustomerAggregateBuilder {
	b.abKey = abKey
	return b
}

func (b *CustomerAggregateBuilder) AddTmcClientNumber(tmcClientNumber string) *CustomerAggregateBuilder {
	b.tmcClientNumber = tmcClientNumber
	return b
}

func (b *CustomerAggregateBuilder) AddAccountNumber(accountNumber string) *CustomerAggregateBuilder {
	b.accountNumber = accountNumber
	return b
}

func (b *CustomerAggregateBuilder) AddState(state string) *CustomerAggregateBuilder {
	b.state = state
	return b
}

func (b *CustomerAggregateBuilder) AddInvoices(invoices []string) *CustomerAggregateBuilder {
	b.invoices = invoices
	return b
}

func (b *CustomerAggregateBuilder) AddPayments(payments []string) *CustomerAggregateBuilder {
	b.payments = payments
	return b
}

func (b *CustomerAggregateBuilder) validate() error {
	if b.name == "" {
		return  errors.New("name is required")
	}
	if b.alias == "" {
		return errors.New("alias is required")
	}
	if b.abKey == "" {
		return errors.New("abKey is required")
	}
	if b.tmcClientNumber == "" {
		return errors.New("tmcClientNumber is required")
	}
	if b.accountNumber == "" {
		return errors.New("accountNumber is required")
	}
	if b.state == "" {
		return errors.New("state is required")
	}
	return nil
}

func (b *CustomerAggregateBuilder) Build() *domain.CustomerAggregate {
	customer := new(domain.CustomerAggregate)
	return &customer{
		Name:            b.name,
		Alias:           b.alias,
		AbKey:           b.abKey,
		TmcClientNumber: b.tmcClientNumber,
		AccountNumber:   b.accountNumber,
		State:           b.state,
		CountryId:       b.countryId,
		Tag:             b.tag,
		Invoices:        b.invoices,
		Payments:        b.payments,
	}
}

func (b *CustomerAggregateBuilder) CreateFromDTO(dto dto.CreateCustomerDTO) *CustomerAggregateBuilder {
	b.AddName(dto.Customer_name)
	b.AddAlias(dto.Alias)
	b.AddAbKey(utils.GenerateRandomString(15)) // Set the implicit AbKey value
	b.AddTmcClientNumber(dto.Tmc_client_number)
	b.AddAccountNumber(dto.Account_number)
	b.AddState(dto.State)
	return b
}
