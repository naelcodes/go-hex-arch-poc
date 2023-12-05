package paymentDomain

import (
	"github.com/naelcodes/ab-backend/pkg/types"
	"github.com/naelcodes/ab-backend/pkg/utils"
)

type PaymentBuilder struct {
	payment *Payment
}

func NewPaymentBuilder() *PaymentBuilder {
	builder := new(PaymentBuilder)
	builder.payment = new(Payment)

	return builder
}

func (builder *PaymentBuilder) SetId(id types.EID) *PaymentBuilder {
	builder.payment.Id = id
	return builder
}

func (builder *PaymentBuilder) SetPaymentNumber(paymentCount int) *PaymentBuilder {
	builder.payment.PaymentNumber = utils.GenerateCode("pr", paymentCount)
	return builder
}

func (builder *PaymentBuilder) SetPaymentDate() *PaymentBuilder {
	builder.payment.PaymentDate = utils.GetCurrentDate()
	return builder
}

func (builder *PaymentBuilder) SetPaymentMode(paymentMode string) *PaymentBuilder {
	builder.payment.PaymentMode = paymentMode
	return builder
}

func (builder *PaymentBuilder) SetAmount(amount float64) *PaymentBuilder {
	builder.payment.Amount = amount
	return builder
}

func (builder *PaymentBuilder) SetBalance(balance float64) *PaymentBuilder {
	builder.payment.Balance = balance
	return builder
}

func (builder *PaymentBuilder) SetUsedAmount(usedAmount float64) *PaymentBuilder {
	builder.payment.UsedAmount = usedAmount
	return builder
}

func (builder *PaymentBuilder) SetStatus(status string) *PaymentBuilder {
	builder.payment.Status = status
	return builder
}

func (builder *PaymentBuilder) SetIdCustomer(idCustomer types.EID) *PaymentBuilder {
	builder.payment.IdCustomer = idCustomer
	return builder
}

func (builder *PaymentBuilder) Build() *Payment {
	return builder.payment
}
