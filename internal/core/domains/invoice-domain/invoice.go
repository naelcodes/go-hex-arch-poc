package invoiceDomain

import (
	"github.com/naelcodes/ab-backend/pkg/types"
)

type Imputation struct {
	types.BaseEntity
	IdInvoice         types.EID
	IdPaymentReceived types.EID
	AmountApply       float64
	PaymentAmount     float64
	InvoiceAmount     float64
}

type TravelItem struct {
	types.BaseEntity
	TotalPrice        float64
	Itinerary         string
	TravelerName      string
	TicketNumber      string
	ConjunctionNumber int
	Status            string
}

type InvoiceAggregate struct {
	types.BaseEntity
	CreationDate  string
	InvoiceNumber string
	Status        string
	DueDate       string
	Amount        float64
	Balance       float64
	Credit_apply  float64
	IdCustomer    types.EID
	Imputations   []Imputation
	TravelItems   []TravelItem
}

type InvoiceDomain struct {
}
