package application

import "github.com/naelcodes/ab-backend/internal/modules/payments/ports"

type PaymentApplication struct {
	repository ports.IPaymentRepository
}

func (application *PaymentApplication) Init(repository ports.IPaymentRepository) {
	application.repository = repository
}
