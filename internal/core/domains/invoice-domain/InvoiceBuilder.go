package invoiceDomain

import (
	"errors"
	"fmt"

	CustomErrors "github.com/naelcodes/ab-backend/pkg/errors"
	"github.com/naelcodes/ab-backend/pkg/types"
	"github.com/naelcodes/ab-backend/pkg/utils"
)

type InvoiceBuilder struct {
	invoice *Invoice
	errors  error
}

func NewInvoiceBuilder() *InvoiceBuilder {
	builder := new(InvoiceBuilder)
	builder.invoice = new(Invoice)
	return builder
}

func (builder *InvoiceBuilder) SetId(id types.EID) *InvoiceBuilder {
	builder.invoice.Id = id
	return builder
}

func (builder *InvoiceBuilder) SetCreationDate(creationDate string) *InvoiceBuilder {
	builder.invoice.CreationDate = creationDate
	return builder
}

func (builder *InvoiceBuilder) SetInvoiceNumber(invoiceCount int) *InvoiceBuilder {
	builder.invoice.InvoiceNumber = utils.GenerateCode("INV", invoiceCount+1)
	return builder
}

func (builder *InvoiceBuilder) SetDueDate(dueDate string) *InvoiceBuilder {
	builder.invoice.DueDate = dueDate
	return builder
}

func (builder *InvoiceBuilder) SetIdCustomer(idCustomer types.EID) *InvoiceBuilder {
	builder.invoice.IdCustomer = idCustomer
	return builder
}

func (builder *InvoiceBuilder) SetAmount(amount float64) *InvoiceBuilder {
	builder.invoice.Amount = amount
	return builder
}

func (builder *InvoiceBuilder) SetCreditApply(creditApply float64) *InvoiceBuilder {
	builder.invoice.Credit_apply = creditApply
	return builder
}

func (builder *InvoiceBuilder) SetBalance(balance float64) *InvoiceBuilder {
	builder.invoice.Balance = balance
	return builder
}

func (builder *InvoiceBuilder) SetStatus(status string) *InvoiceBuilder {
	builder.invoice.Status = status
	return builder
}

func (builder *InvoiceBuilder) SetTravelItemsId(travelItemsId []int) *InvoiceBuilder {
	builder.invoice.TravelItemsId = travelItemsId
	return builder
}

func (builder *InvoiceBuilder) Validate() error {

	if builder.invoice.Credit_apply > builder.invoice.Amount {
		builder.errors = errors.Join(builder.errors, fmt.Errorf("invoice.credit_apply can't be greater than  invoice.amount"))
	}
	if builder.invoice.Balance < 0 {
		builder.errors = errors.Join(builder.errors, fmt.Errorf("invoice.balance can't be less than 0"))
	}

	if builder.invoice.Balance != (builder.invoice.Amount - builder.invoice.Credit_apply) {
		builder.errors = errors.Join(builder.errors, fmt.Errorf("invoice.balance is not equal to invoice.amount - invoice.credit_apply"))
	}

	if builder.invoice.Credit_apply < 0 {
		builder.errors = errors.Join(builder.errors, fmt.Errorf("invoice.credit_apply can't be less than 0"))
	}

	err := builder.invoice.CheckDates()
	if err != nil {
		builder.errors = errors.Join(builder.errors, err)
	}

	if builder.errors != nil {
		return CustomErrors.DomainError(builder.errors)
	}

	return nil
}

func (builder *InvoiceBuilder) Build() *Invoice {
	return builder.invoice
}
