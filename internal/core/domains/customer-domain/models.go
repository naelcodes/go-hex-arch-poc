package customerDomain

import (
	"github.com/naelcodes/ab-backend/pkg/types"
)

type CustomerAggregate struct {
	types.BaseEntity
	CustomerName    string
	Alias           string
	AbKey           string
	TmcClientNumber string
	AccountNumber   string
	State           string
}
