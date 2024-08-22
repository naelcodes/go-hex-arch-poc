package imputationDomain

import (
	"github.com/naelcodes/ab-backend/pkg/types"
)

type ImputationBuilder struct {
	imputation *Imputation
}

func NewImputationBuilder() *ImputationBuilder {
	builder := new(ImputationBuilder)
	builder.imputation = new(Imputation)
	return builder
}

func (builder *ImputationBuilder) SetId(id types.EID) *ImputationBuilder {
	builder.imputation.Id = id
	return builder
}

func (builder *ImputationBuilder) SetIdInvoice(idInvoice types.EID) *ImputationBuilder {
	builder.imputation.IdInvoice = idInvoice
	return builder
}

func (builder *ImputationBuilder) SetIdPayment(idPayment types.EID) *ImputationBuilder {
	builder.imputation.IdPayment = idPayment
	return builder
}

func (builder *ImputationBuilder) SetAmountApplied(amountApply float64) *ImputationBuilder {
	builder.imputation.AmountApplied = amountApply
	return builder
}

func (builder *ImputationBuilder) Build() *Imputation {
	return builder.imputation
}
