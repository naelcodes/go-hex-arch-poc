package dto

import (
	"errors"

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
	PaymentDate   string  `json:"paymentDate"`
}

type InvoiceImputationDTO struct {
	IdPayment     int     `json:"idPayment"`
	AmountApplied float64 `json:"amountApplied"`
}

func (i InvoiceImputationDTO) Validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(&i.IdPayment, validation.Required),
		validation.Field(&i.AmountApplied, validation.By(func(value any) error {
			floatValue, ok := value.(float64)
			if !ok {
				return errors.New("validation error : imputation amount must be a numeric value")
			}

			if floatValue < 0 {
				return errors.New("validation error : imputation amount must be greater than  zero")
			}
			return nil
		})),
	)
}

type GetAllImputationsDTO types.GetAllDTO[[]*GetInvoiceImputationDTO]
