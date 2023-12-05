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

type ImputationDetails struct {
	IdPayment     int     `json:"idPayment"`
	AmountApplied float64 `json:"amountApplied"`
}

func (i ImputationDetails) Validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(&i.IdPayment, validation.Required),
		validation.Field(&i.AmountApplied, validation.Required),
	)
}

type InvoiceImputationsDTO struct {
	IdInvoice   int                 `json:"idInvoice"`
	Imputations []ImputationDetails `json:"imputations"`
}

func (i InvoiceImputationsDTO) Validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(&i.IdInvoice, validation.Required),
		validation.Field(&i.Imputations, validation.Required),
	)
}

type GetAllImputationsDTO types.GetAllDTO[[]*GetInvoiceImputationDTO]
