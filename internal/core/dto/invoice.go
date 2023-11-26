package dto

import "github.com/naelcodes/ab-backend/pkg/types"

type GetInvoiceDTO struct {
	Id            uint    `json:"id"`
	InvoiceNumber string  `json:"invoiceNumber"`
	CreationDate  string  `json:"creationDate"`
	DueDate       string  `json:"dueDate"`
	Amount        float64 `json:"amount"`
	Balance       float64 `json:"balance"`
	Credit_apply  float64 `json:"credit_apply"`
	Status        string  `json:"status"`
}

type GetAllInvoiceDTO types.GetAllDTO[[]*GetInvoiceDTO]
