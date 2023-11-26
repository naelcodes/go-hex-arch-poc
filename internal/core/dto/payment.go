package dto

import "github.com/naelcodes/ab-backend/pkg/types"

type GetPaymentDTO struct {
	Id            uint    `json:"id"`
	PaymentNumber string  `json:"paymentNumber"`
	PaymentDate   string  `json:"paymentDate"`
	PaymentMode   string  `json:"paymentMode"`
	Amount        float64 `json:"amount"`
	Balance       float64 `json:"balance"`
	UsedAmount    float64 `json:"usedAmount"`
	Status        string  `json:"status"`
}

type GetAllPaymentsDTO types.GetAllDTO[[]*GetPaymentDTO]
