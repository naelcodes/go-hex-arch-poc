package domain

import "github.com/naelcodes/ab-backend/internal/common"

type ImputationAggregate struct {
	Id            uint
	InvoiceId     uint
	PaymentId     uint
	AmountApplied common.Money
	PaymentAmount common.Money
	InvoiceAmount common.Money
}

func (ImputationAggregate) ApplyImputation(amount common.Money) {}
