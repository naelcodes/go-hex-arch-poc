package customerDomain

import (
	"errors"

	invoiceDomain "github.com/naelcodes/ab-backend/internal/core/domains/invoice-domain"
	paymentDomain "github.com/naelcodes/ab-backend/internal/core/domains/payment-domain"
	CustomErrors "github.com/naelcodes/ab-backend/pkg/errors"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type CustomerDomainService struct {
	CustomerRepository ICustomerRepository
	InvoiceRepository  invoiceDomain.IInvoiceRepository
	PaymentRepository  paymentDomain.IPaymentRepository
}

func NewCustomerDomainService(customerRepository ICustomerRepository, invoiceRepository invoiceDomain.IInvoiceRepository, paymentRepository paymentDomain.IPaymentRepository) *CustomerDomainService {
	return &CustomerDomainService{
		CustomerRepository: customerRepository,
		InvoiceRepository:  invoiceRepository,
		PaymentRepository:  paymentRepository,
	}
}

func (service *CustomerDomainService) RemoveCustomer(IdCustomer types.EID) error {

	customerInvoiceCount, err := service.InvoiceRepository.CountByCustomerId(types.EID(IdCustomer))

	if err != nil {
		return CustomErrors.RepositoryError(err)
	}

	if *customerInvoiceCount > 0 {
		return CustomErrors.DomainError(errors.New("cannot remove customer  with invoices"))
	}

	customerPaymentCount, err := service.PaymentRepository.CountByCustomerID(types.EID(IdCustomer))

	if err != nil {
		return CustomErrors.RepositoryError(err)
	}

	if *customerPaymentCount > 0 {
		return CustomErrors.DomainError(errors.New("cannot remove customer  with payments"))
	}

	err = service.CustomerRepository.Delete(types.EID(IdCustomer))
	if err != nil {
		return CustomErrors.RepositoryError(err)
	}

	return nil
}
