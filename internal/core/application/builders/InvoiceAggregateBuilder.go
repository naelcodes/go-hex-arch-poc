package builder

import (
	invoiceDomain "github.com/naelcodes/ab-backend/internal/core/domains/invoice-domain"
	"github.com/naelcodes/ab-backend/pkg/types"
	"github.com/naelcodes/ab-backend/pkg/utils"
)

type InvoiceAggregateBuilder struct {
	invoiceAggregate *invoiceDomain.InvoiceAggregate
}

func (i *InvoiceAggregateBuilder) SetCreationDate(creationDate string) *InvoiceAggregateBuilder {
	i.invoiceAggregate.CreationDate = creationDate
	return i
}

func (i *InvoiceAggregateBuilder) SetInvoiceNumber(invoiceCount int) *InvoiceAggregateBuilder {

	i.invoiceAggregate.InvoiceNumber = utils.GenerateCode("INV", invoiceCount+1)
	return i
}

// func (i *InvoiceAggregate) AddTravelItem(TravelItem TravelItem) *InvoiceAggregate {
// 	if len(i.Imputations) > 0 {

// 	}
// 	i.TravelItems = append(i.TravelItems, TravelItem)
// 	i.Amount += TravelItem.TotalPrice
// 	return i
// }

// func (i *InvoiceAggregate) RemoveTravelItem(TravelItem TravelItem) *InvoiceAggregate {
// 	if len(i.Imputations) > 0 {
// 	}
// 	i.Amount -= TravelItem.TotalPrice
// 	return i
// }

func (i *InvoiceAggregateBuilder) SetDueDate(dueDate string) *InvoiceAggregateBuilder {
	i.invoiceAggregate.DueDate = dueDate
	return i
}

func (i *InvoiceAggregateBuilder) SetIdCustomer(idCustomer types.EID) *InvoiceAggregateBuilder {
	i.invoiceAggregate.IdCustomer = idCustomer
	return i
}

func (i *InvoiceAggregateBuilder) Build() *invoiceDomain.InvoiceAggregate {
	return i.invoiceAggregate
}
