package domain

import "github.com/naelcodes/ab-backend/internal/common/base"

type CustomerEntity struct {
	base.IBaseEntity
	base.BaseEntity

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
