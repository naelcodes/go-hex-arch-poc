package postgres

import (
	"github.com/naelcodes/ab-backend/ent"
	"github.com/naelcodes/ab-backend/internal/core/dto"
)

func CustomerModelToDTO(customer *ent.Customer) *dto.GetCustomerDTO {
	customerDTO := new(dto.GetCustomerDTO)

	customerDTO.Id = customer.ID
	customerDTO.Customer_name = customer.CustomerName
	customerDTO.State = customer.State
	customerDTO.Account_number = customer.AccountNumber
	customerDTO.Alias = customer.Alias
	customerDTO.Ab_key = customer.AbKey
	customerDTO.Tmc_client_number = customer.TmcClientNumber

	return customerDTO
}

func CustomerModelListToDTOList(customers []*ent.Customer) []*dto.GetCustomerDTO {
	customerDTOList := make([]*dto.GetCustomerDTO, 0)

	for _, customer := range customers {
		customerDTOList = append(customerDTOList, CustomerModelToDTO(customer))
	}

	return customerDTOList
}

func PaymentModelToDTO(payment *ent.Payment, embedCustomer bool) *dto.GetPaymentDTO {
	paymentDTO := new(dto.GetPaymentDTO)
	paymentDTO.Id = payment.ID
	paymentDTO.Amount = payment.Amount
	paymentDTO.PaymentNumber = payment.Number
	paymentDTO.PaymentDate = payment.Date
	paymentDTO.PaymentMode = string(payment.Fop)
	paymentDTO.Balance = payment.Balance
	paymentDTO.UsedAmount = payment.UsedAmount
	paymentDTO.Status = string(payment.Status)

	if embedCustomer {
		paymentDTO.IdCUstomer = nil
		paymentDTO.Customer = CustomerModelToDTO(payment.Edges.Customer)
	} else {
		paymentDTO.IdCUstomer = &payment.Edges.Customer.ID
	}

	return paymentDTO
}

func PaymentModelListToDTOList(payments []*ent.Payment, embedCustomer bool) []*dto.GetPaymentDTO {
	paymentDTOList := make([]*dto.GetPaymentDTO, 0)

	for _, payment := range payments {
		paymentDTOList = append(paymentDTOList, PaymentModelToDTO(payment, embedCustomer))
	}

	return paymentDTOList
}
