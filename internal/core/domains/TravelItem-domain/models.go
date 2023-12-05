package travelItemDomain

import "github.com/naelcodes/ab-backend/pkg/types"

type TravelItem struct {
	types.BaseEntity
	IdInvoice         *types.EID
	TotalPrice        float64
	Itinerary         *string
	TravelerName      *string
	TicketNumber      *string
	ConjunctionNumber *int
	Status            *string
}

// func (t *TravelItem) AddToInvoice(idInvoice types.EID) {
// 	t.IdInvoice = &idInvoice
// 	t.Status = new(string)
// 	*t.Status = "invoiced"

// }

// func (t *TravelItem) RemoveFromInvoice() {
// 	t.IdInvoice = nil
// 	*t.Status = "pending"
// }
