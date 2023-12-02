package builder

import (
	paymentDomain "github.com/naelcodes/ab-backend/internal/core/domains/payment-domain"
	"github.com/naelcodes/ab-backend/pkg/types"
	"github.com/naelcodes/ab-backend/pkg/utils"
)

type PaymentAggregateBuilder struct {
	paymentAggregate *paymentDomain.PaymentAggregate
}

func NewPaymentAggregateBuilder() *PaymentAggregateBuilder {
	builder := new(PaymentAggregateBuilder)
	builder.paymentAggregate = new(paymentDomain.PaymentAggregate)

	return builder
}

func (builder *PaymentAggregateBuilder) SetId(id types.EID) *PaymentAggregateBuilder {
	builder.paymentAggregate.Id = id
	return builder
}

func (builder *PaymentAggregateBuilder) SetPaymentNumber(paymentCount int) *PaymentAggregateBuilder {
	builder.paymentAggregate.PaymentNumber = utils.GenerateCode("pr", paymentCount)
	return builder
}

func (builder *PaymentAggregateBuilder) SetPaymentDate() *PaymentAggregateBuilder {
	builder.paymentAggregate.PaymentDate = utils.GetCurrentDate()
	return builder
}

func (builder *PaymentAggregateBuilder) SetPaymentMode(paymentMode string) *PaymentAggregateBuilder {
	builder.paymentAggregate.PaymentMode = paymentMode
	return builder
}

func (builder *PaymentAggregateBuilder) SetAmount(amount float64) *PaymentAggregateBuilder {
	builder.paymentAggregate.Amount = amount
	return builder
}

func (builder *PaymentAggregateBuilder) SetBalance(balance float64) *PaymentAggregateBuilder {
	builder.paymentAggregate.Balance = balance
	return builder
}

func (builder *PaymentAggregateBuilder) SetUsedAmount(usedAmount float64) *PaymentAggregateBuilder {
	builder.paymentAggregate.UsedAmount = usedAmount
	return builder
}

func (builder *PaymentAggregateBuilder) SetStatus(status string) *PaymentAggregateBuilder {
	builder.paymentAggregate.Status = status
	return builder
}

func (builder *PaymentAggregateBuilder) SetIdCustomer(idCustomer types.EID) *PaymentAggregateBuilder {
	builder.paymentAggregate.IdCustomer = idCustomer
	return builder
}

func (builder *PaymentAggregateBuilder) Build() *paymentDomain.PaymentAggregate {
	return builder.paymentAggregate
}
