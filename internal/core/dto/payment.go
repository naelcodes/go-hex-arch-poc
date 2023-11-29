package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type GetPaymentDTO struct {
	Id            int             `json:"id"`
	PaymentNumber string          `json:"paymentNumber"`
	PaymentDate   string          `json:"paymentDate"`
	PaymentMode   string          `json:"paymentMode"`
	Amount        float64         `json:"amount"`
	Balance       float64         `json:"balance"`
	UsedAmount    float64         `json:"usedAmount"`
	Status        string          `json:"status"`
	IdCUstomer    *int            `json:"idCustomer,omitempty"`
	Customer      *GetCustomerDTO `json:"customer,omitempty"`
}

type GetAllPaymentsDTO types.GetAllDTO[[]*GetPaymentDTO]

type CreatePaymentDTO struct {
	IdCustomer  int     `json:"idCustomer"`
	Amount      float64 `json:"amount"`
	PaymentMode string  `json:"paymentMode"`
}

func (c CreatePaymentDTO) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.IdCustomer, validation.Required),
		validation.Field(&c.Amount, validation.Required),
		validation.Field(&c.PaymentMode, validation.Required),
	)
}

type UpdatePaymentDTO struct {
	Id          int      `json:"id"`
	IdCustomer  *int     `json:"idCustomer,omitempty"`
	Amount      *float64 `json:"amount,omitempty"`
	PaymentMode *string  `json:"paymentMode,omitempty"`
}

func (u UpdatePaymentDTO) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Id, validation.Required),
		validation.Field(&u.IdCustomer, validation.NilOrNotEmpty),
		validation.Field(&u.Amount, validation.NilOrNotEmpty),
		validation.Field(&u.PaymentMode, validation.NilOrNotEmpty),
	)
}
