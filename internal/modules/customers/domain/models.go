package domain

type CustomerAggregate struct {
	Id              int
	Name            string
	Alias           string
	AbKey           string
	TmcClientNumber string
	AccountNumber   string
	State           string
	CountryId       int
	Tag             string
}

type CountryVO struct {
	Id   uint
	Name string
	Code string
}
