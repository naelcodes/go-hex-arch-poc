package parser

import (
	"github.com/naelcodes/ab-backend/internal/ent"
	"github.com/naelcodes/ab-backend/internal/modules/customers/domain"
)

func CustomerModelListToAggregateList(customers []*ent.Customer) []*domain.CustomerAggregate {
	customerAggregateList := make([]*domain.CustomerAggregate, 0)

	for _, customer := range customers {
		customerAggregateList = append(customerAggregateList, CustomerModelToAggregate(customer))
	}

	return customerAggregateList
}

func CustomerModelToAggregate(customer *ent.Customer) *domain.CustomerAggregate {
	customerAggregate := new(domain.CustomerAggregate)
	customerAggregate.Id = customer.ID
	customerAggregate.Name = customer.CustomerName
	customerAggregate.AccountNumber = customer.AccountNumber
	customerAggregate.TmcClientNumber = customer.TmcClientNumber
	customerAggregate.AbKey = customer.AbKey
	customerAggregate.CountryId = customer.IDCountry
	customerAggregate.Alias = customer.Alias
	return customerAggregate
}
