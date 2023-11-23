package dto

import "github.com/naelcodes/ab-backend/internal/common"

type GetCountriesDTO struct {
	Id   uint   `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type GetCustomerDTO struct {
	Id                uint   ` json:"id"`
	Customer_name     string `json:"customerName"`
	State             string `json:"state"`
	Account_number    string `json:"accountNumber"`
	Id_country        int    `json:"idCountry"`
	Alias             string `json:"alias"`
	Ab_key            string `json:"abKey"`
	Tmc_client_number string `json:"tmcClientNumber"`
}

type CustomerPayment struct {
	Id            uint    `json:"id"`
	PaymentNumber string  `json:"paymentNumber"`
	PaymentDate   string  `json:"paymentDate"`
	PaymentMode   string  `json:"paymentMode"`
	Amount        float64 `json:"amount"`
	Balance       float64 `json:"balance"`
	UsedAmount    float64 `json:"usedAmount"`
	Status        string  `json:"status"`
}

type CustomerInvoice struct {
	Id            uint    `json:"id"`
	InvoiceNumber string  `json:"invoiceNumber"`
	CreationDate  string  `json:"creationDate"`
	DueDate       string  `json:"dueDate"`
	Amount        float64 `json:"amount"`
	Balance       float64 `json:"balance"`
	Credit_apply  float64 `json:"credit_apply"`
	Status        string  `json:"status"`
}

type GetAllCustomersDTO common.GetAllDTO[[]*GetCustomerDTO]
type GetCustomerOpenPaymentsDTO common.GetAllDTO[[]*CustomerPayment]
type GetCustomerUnpaidInvoicesDTO common.GetAllDTO[[]*CustomerInvoice]
