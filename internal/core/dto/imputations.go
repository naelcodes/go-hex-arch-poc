package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type GetInvoiceImputationDTO struct {
	IdInvoice   int                    `json:"idInvoice"`
	Imputations []GetImputationDetails `json:"imputations"`
}

type GetImputationDetails struct {
	Payment       PaymentDetails `json:"payment"`
	AmountApplied float64        `json:"amountApplied"`
}

type PaymentDetails struct {
	Id            int     `json:"id"`
	PaymentNumber string  `json:"paymentNumber"`
	Amount        float64 `json:"amount"`
	Balance       float64 `json:"balance"`
}

type CreateOrUpdateImputationDetails struct {
	IdPayment     int     `json:"idPayment"`
	AmountApplied float64 `json:"amountApplied"`
}

func (c CreateOrUpdateImputationDetails) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.IdPayment, validation.Required),
		validation.Field(&c.AmountApplied, validation.Required),
	)
}

type CreateOrUpdateInvoiceImputationDTO struct {
	IdInvoice   int                               `json:"idInvoice"`
	Imputations []CreateOrUpdateImputationDetails `json:"imputations"`
}

func (c CreateOrUpdateInvoiceImputationDTO) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.IdInvoice, validation.Required),
		validation.Field(&c.Imputations, validation.Required),
	)
}

type GetAllImputationsDTO types.GetAllDTO[[]*GetInvoiceImputationDTO]
