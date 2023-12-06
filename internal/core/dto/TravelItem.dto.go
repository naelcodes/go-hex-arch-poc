package dto

import (
	"errors"

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
		validation.Field(&t.TotalPrice, validation.Required, validation.By(func(value any) error {
			floatValue, ok := value.(float64)
			if !ok {
				return errors.New("validation error : Total price must be a numeric value")
			}

			if floatValue < 0 {
				return errors.New("validation error :Total price must be greater than or equal zero")
			}
			return nil
		})),
	)

}

type GetAllTravelItemDTO types.GetAllDTO[[]*TravelItemDTO]
