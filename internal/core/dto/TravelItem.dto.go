package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type TravelItemDTO struct {
	Id           int     `json:"id"`
	Itinerary    *string `json:"itinerary,omitempty"`
	TravelerName *string `json:"travelerName,omitempty"`
	TicketNumber *string `json:"ticketNumber,omitempty"`
	TotalPrice   float64 `json:"totalPrice"`
}

func (t TravelItemDTO) Validate() error {
	return validation.ValidateStruct(&t,
		validation.Field(&t.Id, validation.Required),
		validation.Field(&t.TotalPrice, validation.Required),
	)
}

type GetAllTravelItemDTO types.GetAllDTO[[]*TravelItemDTO]
