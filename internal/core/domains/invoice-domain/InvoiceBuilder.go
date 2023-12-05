package invoiceDomain

import (
	"github.com/naelcodes/ab-backend/pkg/types"
	"github.com/naelcodes/ab-backend/pkg/utils"
)

type InvoiceBuilder struct {
	invoice *Invoice
}

func NewInvoiceBuilder() *InvoiceBuilder {
	builder := new(InvoiceBuilder)
	builder.invoice = new(Invoice)
	return builder
}

func (i *InvoiceBuilder) SetId(id types.EID) *InvoiceBuilder {
	i.invoice.Id = id
	return i
}

func (i *InvoiceBuilder) SetCreationDate(creationDate string) *InvoiceBuilder {
	i.invoice.CreationDate = creationDate
	return i
}

func (i *InvoiceBuilder) SetInvoiceNumber(invoiceCount int) *InvoiceBuilder {
	i.invoice.InvoiceNumber = utils.GenerateCode("INV", invoiceCount+1)
	return i
}

func (i *InvoiceBuilder) SetDueDate(dueDate string) *InvoiceBuilder {
	i.invoice.DueDate = dueDate
	return i
}

func (i *InvoiceBuilder) SetIdCustomer(idCustomer types.EID) *InvoiceBuilder {
	i.invoice.IdCustomer = idCustomer
	return i
}

func (i *InvoiceBuilder) SetAmount(amount float64) *InvoiceBuilder {
	i.invoice.Amount = amount
	return i
}

func (i *InvoiceBuilder) SetCreditApply(creditApply float64) *InvoiceBuilder {
	i.invoice.Credit_apply = creditApply
	return i
}

func (i *InvoiceBuilder) SetBalance(balance float64) *InvoiceBuilder {
	i.invoice.Balance = balance
	return i
}

func (i *InvoiceBuilder) SetStatus(status string) *InvoiceBuilder {
	i.invoice.Status = status
	return i
}

func (i *InvoiceBuilder) Build() *Invoice {
	return i.invoice
}
