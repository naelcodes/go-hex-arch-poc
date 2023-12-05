package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type CreateCustomerDTO struct {
	CustomerName    string `json:"customerName"`
	State           string `json:"state"`
	AccountNumber   string `json:"accountNumber"`
	Alias           string `json:"alias"`
	TmcClientNumber string `json:"tmcClientNumber"`
}

func (c CreateCustomerDTO) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.CustomerName, validation.Required),
		validation.Field(&c.State, validation.Required),
		validation.Field(&c.AccountNumber, validation.Required),
		validation.Field(&c.Alias, validation.Required),
		validation.Field(&c.TmcClientNumber, validation.Required),
	)
}

type UpdateCustomerDTO struct {
	Customer_name     *string `json:"customerName,omitempty"`
	State             *string `json:"state,omitempty"`
	Account_number    *string `json:"accountNumber,omitempty"`
	Alias             *string `json:"alias,omitempty"`
	Tmc_client_number *string `json:"tmcClientNumber,omitempty"`
}

func (u UpdateCustomerDTO) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Customer_name, validation.NilOrNotEmpty),
		validation.Field(&u.State, validation.NilOrNotEmpty),
		validation.Field(&u.Account_number, validation.NilOrNotEmpty),
		validation.Field(&u.Alias, validation.NilOrNotEmpty),
		validation.Field(&u.Tmc_client_number, validation.NilOrNotEmpty),
	)
}

type GetCustomerDTO struct {
	Id                int    ` json:"id"`
	Customer_name     string `json:"customerName"`
	State             string `json:"state,omitempty"`
	Account_number    string `json:"accountNumber,omitempty"`
	Alias             string `json:"alias,omitempty"`
	Ab_key            string `json:"abKey,omitempty"`
	Tmc_client_number string `json:"tmcClientNumber,omitempty"`
}

type CustomerPayments struct {
	IdCustomer int
	Payments   []*GetPaymentDTO
}
type CustomerInvoice struct {
	IdCustomer int
	Invoices   []*GetInvoiceDTO
}
type GetAllCustomersDTO types.GetAllDTO[[]*GetCustomerDTO]
type GetCustomerPaymentsDTO types.GetAllDTO[*CustomerPayments]
type GetCustomerInvoicesDTO types.GetAllDTO[*CustomerInvoice]
