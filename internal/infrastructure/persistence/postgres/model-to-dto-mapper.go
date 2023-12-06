package postgres

import (
	"github.com/naelcodes/ab-backend/ent"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	"github.com/naelcodes/ab-backend/pkg/utils"
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

func PaymentModelToDTO(payment *ent.Payment, embedCustomer bool, idCustomer *int) *dto.GetPaymentDTO {

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
		if payment.Edges.Customer != nil {
			paymentDTO.IdCUstomer = &payment.Edges.Customer.ID
		} else {
			paymentDTO.IdCUstomer = idCustomer
		}
	}

	return paymentDTO
}

func PaymentModelListToDTOList(payments []*ent.Payment, embedCustomer bool, idCustomer *int) []*dto.GetPaymentDTO {
	paymentDTOList := make([]*dto.GetPaymentDTO, 0)

	for _, payment := range payments {
		paymentDTOList = append(paymentDTOList, PaymentModelToDTO(payment, embedCustomer, idCustomer))
	}

	return paymentDTOList
}

func TravelItemModelToDTO(travelItem *ent.TravelItem) *dto.TravelItemDTO {
	travelItemDTO := new(dto.TravelItemDTO)

	travelItemDTO.Id = travelItem.ID
	travelItemDTO.TotalPrice = travelItem.TotalPrice

	travelItemDTO.Itinerary = new(string)
	*travelItemDTO.Itinerary = travelItem.Itinerary

	travelItemDTO.TravelerName = new(string)
	*travelItemDTO.TravelerName = travelItem.TravelerName

	travelItemDTO.TicketNumber = new(string)
	*travelItemDTO.TicketNumber = travelItem.TicketNumber

	return travelItemDTO
}

func TravelItemModelListToDTOList(travelItems []*ent.TravelItem) []*dto.TravelItemDTO {
	travelItemDTOList := make([]*dto.TravelItemDTO, 0)

	for _, travelItem := range travelItems {
		travelItemDTOList = append(travelItemDTOList, TravelItemModelToDTO(travelItem))
	}

	return travelItemDTOList
}

func InvoiceModelToDTO(invoice *ent.Invoice, embedCustomer bool) *dto.GetInvoiceDTO {
	invoiceDTO := new(dto.GetInvoiceDTO)
	invoiceDTO.Id = invoice.ID
	invoiceDTO.InvoiceNumber = invoice.InvoiceNumber
	invoiceDTO.DueDate = utils.FormatDate(invoice.DueDate)
	invoiceDTO.CreationDate = utils.FormatDate(invoice.CreationDate)
	invoiceDTO.Amount = invoice.Amount
	invoiceDTO.Credit_apply = invoice.CreditApply
	invoiceDTO.Balance = invoice.Balance
	invoiceDTO.Status = string(invoice.Status)

	if embedCustomer {
		invoiceDTO.IdCustomer = nil
		invoiceDTO.Customer = CustomerModelToDTO(invoice.Edges.Customer)
	} else {
		if invoice.Edges.Customer != nil {
			invoiceDTO.IdCustomer = &invoice.Edges.Customer.ID
		}
	}

	for _, travelItem := range invoice.Edges.TravelItems {
		invoiceDTO.TravelItems = append(invoiceDTO.TravelItems, TravelItemModelToDTO(travelItem))
	}
	return invoiceDTO
}

func InvoiceModelListToDTOList(invoices []*ent.Invoice, embedCustomer bool) []*dto.GetInvoiceDTO {
	invoiceDTOList := make([]*dto.GetInvoiceDTO, 0)

	for _, invoice := range invoices {
		invoiceDTOList = append(invoiceDTOList, InvoiceModelToDTO(invoice, embedCustomer))
	}

	return invoiceDTOList
}
