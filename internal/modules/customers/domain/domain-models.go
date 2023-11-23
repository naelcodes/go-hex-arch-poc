package domain

import "github.com/naelcodes/ab-backend/internal/common"

type CustomerAggregate struct {
	common.BaseEntity
	Name            string
	Alias           string
	AbKey           string
	TmcClientNumber string
	AccountNumber   string
	State           string
	CountryId       int
	Tag             string
}

func NewCustomerAggregate(id int, name string, alias string, abkey string, tmcClientNumber string, accountNumber string, state string, countryId int, tag string) *CustomerAggregate {
	c := new(CustomerAggregate)
	c.Id = common.EID(id)
	c.Name = name
	c.Alias = alias
	c.AbKey = abkey
	c.TmcClientNumber = tmcClientNumber
	c.AccountNumber = accountNumber
	c.State = state
	c.CountryId = countryId
	c.Tag = tag

	return c
}
