package builder

import (
	"fmt"

	invoiceDomain "github.com/naelcodes/ab-backend/internal/core/domains/invoice-domain"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type InvoiceAggregateBuilder struct {
	invoiceAggregate *invoiceDomain.InvoiceAggregate
}

func (i *InvoiceAggregateBuilder) SetCreationDate(creationDate string) *InvoiceAggregateBuilder {
	i.invoiceAggregate.CreationDate = creationDate
	return i
}

func (i *InvoiceAggregateBuilder) SetInvoiceNumber(invoiceCount int) *InvoiceAggregateBuilder {

	var invoiceNumber string

	if invoiceCount >= 0 && invoiceCount < 10 {
		invoiceNumber = fmt.Sprintf("INV-000%d", invoiceCount+1)
	} else if invoiceCount >= 10 && invoiceCount < 100 {
		invoiceNumber = fmt.Sprintf("INV-00%d", invoiceCount+1)
	} else if invoiceCount >= 100 && invoiceCount < 1000 {
		invoiceNumber = fmt.Sprintf("INV-0%d", invoiceCount+1)
	} else {
		invoiceNumber = fmt.Sprintf("INV-%d", invoiceCount+1)
	}

	i.invoiceAggregate.InvoiceNumber = invoiceNumber
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

// func (i *InvoiceAggregateBuilder) SetTravelItem(travelItems []TravelItem) *InvoiceAggregateBuilder {
// 	i.invoiceAggregate.TravelItems = travelItems
// 	return i
// }

// func (i *InvoiceAggregateBuilder) SetImputation(imputations []Imputation) *InvoiceAggregateBuilder {
// 	i.invoiceAggregate.Imputations = imputations
// 	return i
// }

func (i *InvoiceAggregateBuilder) Build() *invoiceDomain.InvoiceAggregate {
	return i.invoiceAggregate
}
