package paymentDomain

import (
	"errors"

	imputationDomain "github.com/naelcodes/ab-backend/internal/core/domains/imputation-domain"
	CustomErrors "github.com/naelcodes/ab-backend/pkg/errors"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type PaymentDomainService struct {
	ImputationRepository imputationDomain.IImputationRepository
	PaymentRepository    IPaymentRepository
}

func NewPaymentDomainService(imputationRepository imputationDomain.IImputationRepository, paymentRepository IPaymentRepository) *PaymentDomainService {
	return &PaymentDomainService{
		ImputationRepository: imputationRepository,
		PaymentRepository:    paymentRepository,
	}
}

func (service *PaymentDomainService) RemovePayment(paymentId types.EID) error {

	totalCount, err := service.ImputationRepository.CountByPaymentId(paymentId)

	if err != nil {
		return err
	}

	if *totalCount > 0 {
		return CustomErrors.DomainError(errors.New("cannot remove payment with imputations"))
	}

	err = service.PaymentRepository.Delete(paymentId)

	if err != nil {
		return err
	}

	return nil
}
