package application

import "github.com/naelcodes/ab-backend/internal/modules/invoices/ports"

type InvoiceApplication struct {
	repository ports.IInvoiceRepository
}

func (application *InvoiceApplication) Init(invoiceRepository ports.IInvoiceRepository) {
	application.repository = invoiceRepository
}
