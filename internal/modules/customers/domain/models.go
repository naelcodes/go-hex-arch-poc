package domain

type CustomerAggregate struct {
	Id              uint
	Name            string
	Alias           string
	AbKey           string
	TmcClientNumber string
	AccountNumber   string

	Country CountryVO
	Slug    uint64
	Tag     string
}

type CountryVO struct {
	Id   uint
	Name string
	Code string
}
