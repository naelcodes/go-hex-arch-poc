package domain

import (
	"time"

	"github.com/naelcodes/ab-backend/internal/common"
)

type TravelItemEntity struct{}

type InvoiceAggregate struct {
	id             uint
	invoice_number string
	status         string
	due_date       time.Time
	amount         common.Money
	balance        common.Money
	creditApplied  common.Money
	slug           int64
}

func (InvoiceAggregate) New() {

}
