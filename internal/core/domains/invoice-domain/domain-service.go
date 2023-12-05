package invoiceDomain

import (
	"errors"

	travelItemDomain "github.com/naelcodes/ab-backend/internal/core/domains/TravelItem-domain"
	imputationDomain "github.com/naelcodes/ab-backend/internal/core/domains/imputation-domain"
	paymentDomain "github.com/naelcodes/ab-backend/internal/core/domains/payment-domain"
	CustomErrors "github.com/naelcodes/ab-backend/pkg/errors"
	"github.com/naelcodes/ab-backend/pkg/transactionManager"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type InvoiceDomainService struct {
	travelItemRepository travelItemDomain.ITravelItemRepository
	imputationRepository imputationDomain.IImputationRepository
	paymentRepository    paymentDomain.IPaymentRepository
	invoiceRepository    IInvoiceRepository
	transactionManager   *transactionManager.TransactionManager
}

func NewInvoiceDomainService(
	travelItemRepository travelItemDomain.ITravelItemRepository,
	imputationRepository imputationDomain.IImputationRepository,
	paymentRepository paymentDomain.IPaymentRepository,
	invoiceRepository IInvoiceRepository,
	transactionManager *transactionManager.TransactionManager) *InvoiceDomainService {
	return &InvoiceDomainService{
		travelItemRepository: travelItemRepository,
		imputationRepository: imputationRepository,
		paymentRepository:    paymentRepository,
		invoiceRepository:    invoiceRepository,
		transactionManager:   transactionManager,
	}
}

func (service *InvoiceDomainService) AddTravelItem(invoice *Invoice, travelItems []*travelItemDomain.TravelItem) {

	imputationCount, err := service.imputationRepository.CountByInvoiceId(invoice.Id)

	if err != nil {
		panic(err)
	}

	if *imputationCount > 0 {
		panic(CustomErrors.DomainError(errors.New("cannot add travel item to an invoice  with imputations")))
	}

	travelItemIds := make([]int, 0)

	for _, travelItem := range travelItems {
		travelItemIds = append(travelItemIds, int(travelItem.Id))
		invoice.Amount += travelItem.TotalPrice
	}

	invoiceRepoErr := service.invoiceRepository.Update(service.transactionManager.GetTransaction(), invoice)

	if invoiceRepoErr != nil {
		panic(invoiceRepoErr)
	}

	travelItemRepoErr := service.travelItemRepository.UpdateByInvoiceId(service.transactionManager.GetTransaction(), &invoice.Id, travelItemIds)

	if travelItemRepoErr != nil {
		panic(travelItemRepoErr)
	}
}

func (service *InvoiceDomainService) RemoveTravelItem(invoice Invoice, travelItems []*travelItemDomain.TravelItem) {

	imputationCount, err := service.imputationRepository.CountByInvoiceId(invoice.Id)

	if err != nil {
		panic(err)
	}

	if *imputationCount > 0 {
		panic(CustomErrors.DomainError(errors.New("cannot remove travel item to an invoice  with imputations")))
	}

	travelItemIds := make([]int, 0)

	for _, travelItem := range travelItems {
		travelItemIds = append(travelItemIds, int(travelItem.Id))
		invoice.Amount -= travelItem.TotalPrice
	}

	invoiceRepoErr := service.invoiceRepository.Update(service.transactionManager.GetTransaction(), &invoice)

	if invoiceRepoErr != nil {
		panic(invoiceRepoErr)
	}

	travelItemRepoErr := service.travelItemRepository.UpdateByInvoiceId(service.transactionManager.GetTransaction(), nil, travelItemIds)

	if travelItemRepoErr != nil {
		panic(travelItemRepoErr)
	}
}

func (service *InvoiceDomainService) AddImputation(invoiceId types.EID, imputations []*imputationDomain.Imputation) {

	payments := make([]*paymentDomain.Payment, 0)
	invoiceDTO, RepositoryErr := service.invoiceRepository.GetById(invoiceId)

	if RepositoryErr != nil {
		panic(RepositoryErr)
	}

	invoice := NewInvoiceBuilder().
		SetId(types.EID(invoiceDTO.Id)).
		SetAmount(invoiceDTO.Amount).
		SetBalance(invoiceDTO.Balance).
		SetCreditApply(invoiceDTO.Credit_apply).
		SetStatus(invoiceDTO.Status).
		Build()

	for _, imputation := range imputations {

		paymentDTO, RepositoryErr := service.paymentRepository.GetById(imputation.IdPayment)

		if RepositoryErr != nil {
			panic(RepositoryErr)
		}

		payment := paymentDomain.NewPaymentBuilder().
			SetId(types.EID(paymentDTO.Id)).
			SetStatus(paymentDTO.Status).
			SetAmount(paymentDTO.Amount).
			SetBalance(paymentDTO.Balance).
			SetUsedAmount(paymentDTO.UsedAmount).
			Build()

		// imputation logic on payment (payment domain)
		paymentDomainErr := payment.AllocateAmount(imputation.AmountApplied)

		if paymentDomainErr != nil {
			panic(paymentDomainErr)
		}

		payments = append(payments, payment)

		// imputation logic on invoice (invoice domain)
		invoiceDomainErr := invoice.ApplyImputation(imputation.AmountApplied)
		if invoiceDomainErr != nil {
			panic(invoiceDomainErr)
		}

	}

	service.paymentRepository.SaveAllPaymentsAllocations(service.transactionManager.GetTransaction(), payments)
	service.invoiceRepository.SaveImputation(service.transactionManager.GetTransaction(), invoice)
	service.imputationRepository.SaveAll(service.transactionManager.GetTransaction(), imputations)

}

func (service *InvoiceDomainService) UpdateImputation(invoiceId types.EID, imputations []*imputationDomain.Imputation) {

	payments := make([]*paymentDomain.Payment, 0)
	invoiceDTO, InvoiceRepositoryErr := service.invoiceRepository.GetById(invoiceId)

	if InvoiceRepositoryErr != nil {
		panic(InvoiceRepositoryErr)
	}

	invoice := NewInvoiceBuilder().
		SetId(types.EID(invoiceDTO.Id)).
		SetAmount(invoiceDTO.Amount).
		SetBalance(invoiceDTO.Balance).
		SetCreditApply(invoiceDTO.Credit_apply).
		SetStatus(invoiceDTO.Status).
		Build()

	for _, imputation := range imputations {

		savedImputationDTO, ImputationRepoErr := service.imputationRepository.GetByPaymentAndInvoiceId(imputation.IdPayment, invoiceId)

		if ImputationRepoErr != nil {
			panic(ImputationRepoErr)
		}

		newImputation := imputationDomain.NewImputationBuilder().
			SetId(imputation.Id).
			SetAmountApply(imputation.AmountApplied).
			Build()

		if newImputation.AmountApplied == 0 && savedImputationDTO.AmountApplied > 0 {

			ImputationRepoErr := service.imputationRepository.Delete(service.transactionManager.GetTransaction(), imputation.Id)
			if ImputationRepoErr != nil {
				panic(ImputationRepoErr)
			}

		} else {

			ImputationRepoErr := service.imputationRepository.Update(service.transactionManager.GetTransaction(), newImputation)
			if ImputationRepoErr != nil {
				panic(ImputationRepoErr)
			}
		}

		imputationDiff := newImputation.AmountApplied - savedImputationDTO.AmountApplied

		invoiceDomainErr := invoice.ApplyImputation(imputationDiff)

		if invoiceDomainErr != nil {
			panic(invoiceDomainErr)
		}

		paymentDTO, PaymentRepositoryErr := service.paymentRepository.GetById(imputation.IdPayment)

		if PaymentRepositoryErr != nil {
			panic(PaymentRepositoryErr)
		}

		payment := paymentDomain.NewPaymentBuilder().
			SetId(types.EID(paymentDTO.Id)).
			SetAmount(paymentDTO.Amount).
			SetBalance(paymentDTO.Balance).
			SetUsedAmount(paymentDTO.UsedAmount).
			Build()

		paymentDomainErr := payment.AllocateAmount(imputationDiff)

		if paymentDomainErr != nil {
			panic(paymentDomainErr)
		}
		payments = append(payments, payment)

	}

	service.invoiceRepository.SaveImputation(service.transactionManager.GetTransaction(), invoice)
	service.paymentRepository.SaveAllPaymentsAllocations(service.transactionManager.GetTransaction(), payments)
}

func (service *InvoiceDomainService) VoidInvoice(invoiceId types.EID) error {
	defer service.transactionManager.CatchError()

	return nil
}
