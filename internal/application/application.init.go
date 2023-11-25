package application

import domain "github.com/naelcodes/ab-backend/internal/domain/customer"

type Application struct {
	CustomerRepository domain.ICustomerRepository
	// InvoiceRepository
	// PaymentRepository
}
