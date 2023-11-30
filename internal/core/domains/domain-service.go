package domains

import (
	"errors"

	customerDomain "github.com/naelcodes/ab-backend/internal/core/domains/customer-domain"
	invoiceDomain "github.com/naelcodes/ab-backend/internal/core/domains/invoice-domain"
	paymentDomain "github.com/naelcodes/ab-backend/internal/core/domains/payment-domain"
	CustomErrors "github.com/naelcodes/ab-backend/pkg/errors"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type DomainService struct{}

func (domainService *DomainService) RemoveCustomer(IdCustomer types.EID, CustomerRepository customerDomain.ICustomerRepository, InvoiceRepository invoiceDomain.IInvoiceRepository, PaymentRepository paymentDomain.IPaymentRepository) error {

	customerInvoiceCount, err := InvoiceRepository.CountByCustomerId(types.EID(IdCustomer))

	if err != nil {
		return CustomErrors.RepositoryError(err)
	}

	if customerInvoiceCount > 0 {
		return CustomErrors.DomainError(errors.New("cannot remove customer  with invoices"))
	}

	customerPaymentCount, err := PaymentRepository.CountByCustomerID(types.EID(IdCustomer))

	if err != nil {
		return CustomErrors.RepositoryError(err)
	}

	if customerPaymentCount > 0 {
		return CustomErrors.DomainError(errors.New("cannot remove customer  with payments"))
	}

	err = CustomerRepository.Delete(types.EID(IdCustomer))
	if err != nil {
		return CustomErrors.RepositoryError(err)
	}

	return nil
}
