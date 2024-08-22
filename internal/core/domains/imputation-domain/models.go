package imputationDomain

import "github.com/naelcodes/ab-backend/pkg/types"

type Imputation struct {
	types.BaseEntity
	IdInvoice     types.EID
	IdPayment     types.EID
	AmountApplied float64
}
