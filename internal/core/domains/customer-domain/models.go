package customerDomain

import (
	"github.com/naelcodes/ab-backend/pkg/types"
)

type Customer struct {
	types.BaseEntity
	CustomerName    string
	Alias           string
	AbKey           string
	TmcClientNumber string
	AccountNumber   string
	State           string
}
