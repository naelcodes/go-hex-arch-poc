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
}
