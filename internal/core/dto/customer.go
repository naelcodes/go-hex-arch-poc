package dto

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type CreateCustomerDTO struct {
	Customer_name     string `json:"customerName"`
	State             string `json:"state"`
	Account_number    string `json:"accountNumber"`
	Alias             string `json:"alias"`
	Tmc_client_number string `json:"tmcClientNumber"`
}

func (c CreateCustomerDTO) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Customer_name, validation.Required),
		validation.Field(&c.State, validation.Required),
		validation.Field(&c.Account_number, validation.Required),
		validation.Field(&c.Alias, validation.Required),
		validation.Field(&c.Tmc_client_number, validation.Required),
	)
}

type UpdateCustomerDTO struct {
	Id                int     `json:"id"`
	Customer_name     *string `json:"customerName,omitempty"`
	State             *string `json:"state,omitempty"`
	Account_number    *string `json:"accountNumber,omitempty"`
	Alias             *string `json:"alias,omitempty"`
	Tmc_client_number *string `json:"tmcClientNumber,omitempty"`
}

func (u UpdateCustomerDTO) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Id, validation.Required),
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

type CustomerOpenPayments struct {
	IdCustomer int
	Payments   []GetPaymentDTO
}
type CustomerUnpaidInvoice struct {
	IdCustomer int
	Invoices   []GetInvoiceDTO
}
type GetAllCustomersDTO types.GetAllDTO[[]*GetCustomerDTO]
type GetCustomerOpenPaymentsDTO types.GetAllDTO[*CustomerOpenPayments]
type GetCustomerUnpaidInvoicesDTO types.GetAllDTO[*CustomerUnpaidInvoice]
