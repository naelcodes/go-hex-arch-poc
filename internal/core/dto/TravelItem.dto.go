package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type TravelItemDTO struct {
	Id           int      `json:"id"`
	IdInvoice    *int     `json:"idInvoice,omitempty"`
	Itinerary    *string  `json:"itinerary,omitempty"`
	TravelerName *string  `json:"travelerName,omitempty"`
	TicketNumber *string  `json:"ticketNumber,omitempty"`
	TotalPrice   *float64 `json:"totalPrice,omitempty"`
}

func (t TravelItemDTO) Validate() error {
	return validation.ValidateStruct(&t,
		validation.Field(&t.Id, validation.Required),
	)

}

type GetAllTravelItemDTO types.GetAllDTO[[]*TravelItemDTO]
