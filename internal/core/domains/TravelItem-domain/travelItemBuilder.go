package travelItemDomain

import "github.com/naelcodes/ab-backend/pkg/types"

type TravelItemBuilder struct {
	travelItem *TravelItem
}

func NewTravelItemBuilder() *TravelItemBuilder {
	return &TravelItemBuilder{
		travelItem: &TravelItem{},
	}
}

func (b *TravelItemBuilder) SetId(id types.EID) *TravelItemBuilder {
	b.travelItem.Id = id
	return b
}

func (b *TravelItemBuilder) SetIdInvoice(idInvoice types.EID) *TravelItemBuilder {
	b.travelItem.IdInvoice = &idInvoice
	return b
}

func (b *TravelItemBuilder) SetTotalPrice(totalPrice float64) *TravelItemBuilder {
	b.travelItem.TotalPrice = totalPrice
	return b
}

func (b *TravelItemBuilder) SetItinerary(itinerary *string) *TravelItemBuilder {
	b.travelItem.Itinerary = itinerary
	return b
}

func (b *TravelItemBuilder) SetTravelerName(travelerName *string) *TravelItemBuilder {
	b.travelItem.TravelerName = travelerName
	return b
}

func (b *TravelItemBuilder) SetTicketNumber(ticketNumber *string) *TravelItemBuilder {
	b.travelItem.TicketNumber = ticketNumber
	return b
}

func (b *TravelItemBuilder) SetConjunctionNumber(conjunctionNumber *int) *TravelItemBuilder {
	b.travelItem.ConjunctionNumber = conjunctionNumber
	return b
}

func (b *TravelItemBuilder) SetStatus(status *string) *TravelItemBuilder {
	b.travelItem.Status = status
	return b
}

func (b *TravelItemBuilder) Build() *TravelItem {
	return b.travelItem
}
