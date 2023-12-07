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

	imputationRepository imputationDomain.IImputationRepository,
	paymentRepository paymentDomain.IPaymentRepository,
	invoiceRepository IInvoiceRepository,
	transactionManager *transactionManager.TransactionManager) *InvoiceDomainService {
	return &InvoiceDomainService{
		imputationRepository: imputationRepository,
		paymentRepository:    paymentRepository,
		invoiceRepository:    invoiceRepository,
		transactionManager:   transactionManager,
	}
}

func (service *InvoiceDomainService) ApplyImputation(invoiceId types.EID, imputationDomainModelList []*imputationDomain.Imputation) error {

	savedImputationPaymentIdToAmountMap := make(map[int]float64, 0)
	imputationsToInsert := make([]*imputationDomain.Imputation, 0)
	imputationsToUpdate := make([]*imputationDomain.Imputation, 0)

	invoiceDTO, RepositoryErr := service.invoiceRepository.GetById(invoiceId, nil)

	if RepositoryErr != nil {
		return RepositoryErr
	}

	invoiceBuilder := NewInvoiceBuilder().
		SetId(types.EID(invoiceDTO.Id)).
		SetCreationDate(invoiceDTO.CreationDate).
		SetDueDate(invoiceDTO.DueDate).
		SetIdCustomer(types.EID(*invoiceDTO.IdCustomer)).
		SetAmount(invoiceDTO.Amount).
		SetBalance(invoiceDTO.Balance).
		SetCreditApply(invoiceDTO.Credit_apply).
		SetStatus(invoiceDTO.Status)

	invoiceDomainErr := invoiceBuilder.Validate()
	if invoiceDomainErr != nil {
		return CustomErrors.DomainError(errors.Join(errors.New("error - saved invoice in invalid state"), invoiceDomainErr))
	}

	invoiceDomainModel := invoiceBuilder.Build()

	// Check if imputation already exists and sort them
	for _, imputationDomainModel := range imputationDomainModelList {
		exists, imputationRecord, err := service.imputationRepository.GetByPaymentAndInvoiceId(imputationDomainModel.IdPayment, imputationDomainModel.IdInvoice)

		if err != nil {
			return err
		}

		if *exists {
			imputationDomainModel.Id = types.EID(imputationRecord.ID)
			imputationsToUpdate = append(imputationsToUpdate, imputationDomainModel)
			savedImputationPaymentIdToAmountMap[imputationRecord.Edges.Payment.ID] = imputationRecord.AmountApply
		} else {
			imputationsToInsert = append(imputationsToInsert, imputationDomainModel)
		}
	}

	if len(imputationsToUpdate) > 0 {

		for _, imputationToUpdate := range imputationsToUpdate {

			imputedDiff := imputationToUpdate.AmountApplied - savedImputationPaymentIdToAmountMap[int(imputationToUpdate.IdPayment)]

			// Delete imputation if amount applied is 0
			if imputationToUpdate.AmountApplied == 0 {
				repoErr := service.imputationRepository.Delete(service.transactionManager.GetTransaction(), imputationToUpdate.Id)

				if repoErr != nil {
					return repoErr
				}

			} else {

				//  Update imputation
				repoErr := service.imputationRepository.Update(service.transactionManager.GetTransaction(), imputationToUpdate)

				if repoErr != nil {
					return repoErr
				}
			}

			// apply imputation to payment
			paymentDTO, RepositoryErr := service.paymentRepository.GetById(imputationToUpdate.IdPayment)

			if RepositoryErr != nil {
				return RepositoryErr
			}

			paymentBuilder := paymentDomain.NewPaymentBuilder().
				SetId(types.EID(paymentDTO.Id)).
				SetStatus(paymentDTO.Status).
				SetAmount(paymentDTO.Amount).
				SetBalance(paymentDTO.Balance).
				SetUsedAmount(paymentDTO.UsedAmount)

			domainErr := paymentBuilder.Validate()
			if domainErr != nil {

				return domainErr
			}
			paymentDomainModel := paymentBuilder.Build()
			domainErr = paymentDomainModel.AllocateAmount(imputedDiff)

			if domainErr != nil {
				return domainErr
			}

			repoErr := service.paymentRepository.SavePaymentAllocation(service.transactionManager.GetTransaction(), paymentDomainModel)

			if repoErr != nil {
				return repoErr
			}

			// Case 1.4 - Update invoice domain model
			domainErr = invoiceDomainModel.ApplyImputation(imputedDiff)

			if domainErr != nil {
				return domainErr
			}

		}

	}

	if len(imputationsToInsert) > 0 {

		for _, imputationToInsert := range imputationsToInsert {

			paymentDTO, RepositoryErr := service.paymentRepository.GetById(imputationToInsert.IdPayment)

			if RepositoryErr != nil {
				return RepositoryErr
			}

			paymentBuilder := paymentDomain.NewPaymentBuilder().
				SetId(types.EID(paymentDTO.Id)).
				SetStatus(paymentDTO.Status).
				SetAmount(paymentDTO.Amount).
				SetBalance(paymentDTO.Balance).
				SetUsedAmount(paymentDTO.UsedAmount)

			domainErr := paymentBuilder.Validate()
			if domainErr != nil {
				return domainErr
			}
			paymentDomainModel := paymentBuilder.Build()

			domainErr = paymentDomainModel.AllocateAmount(imputationToInsert.AmountApplied)

			if domainErr != nil {
				return domainErr
			}

			repoErr := service.paymentRepository.SavePaymentAllocation(service.transactionManager.GetTransaction(), paymentDomainModel)

			if repoErr != nil {
				return repoErr
			}

			domainErr = invoiceDomainModel.ApplyImputation(imputationToInsert.AmountApplied)

			if domainErr != nil {
				return domainErr
			}

		}

		RepoErr := service.imputationRepository.SaveAll(service.transactionManager.GetTransaction(), imputationsToInsert)

		if RepoErr != nil {
			return RepoErr
		}
	}

	RepoErr := service.invoiceRepository.SaveImputation(service.transactionManager.GetTransaction(), invoiceDomainModel)
	if RepoErr != nil {
		return RepoErr
	}

	return nil

}

// func (service *InvoiceDomainService) AddTravelItem(invoice *Invoice, travelItems []*travelItemDomain.TravelItem) {

// 	imputationCount, err := service.imputationRepository.CountByInvoiceId(invoice.Id)

// 	if err != nil {
// 		panic(err)
// 	}

// 	if *imputationCount > 0 {
// 		panic(CustomErrors.DomainError(errors.New("cannot add travel item to an invoice  with imputations")))
// 	}

// 	travelItemIds := make([]int, 0)

// 	for _, travelItem := range travelItems {
// 		travelItemIds = append(travelItemIds, int(travelItem.Id))
// 		invoice.Amount += travelItem.TotalPrice
// 	}

// 	invoiceRepoErr := service.invoiceRepository.Update(service.transactionManager.GetTransaction(), invoice)

// 	if invoiceRepoErr != nil {
// 		panic(invoiceRepoErr)
// 	}

// 	travelItemRepoErr := service.travelItemRepository.UpdateByInvoiceId(service.transactionManager.GetTransaction(), &invoice.Id, travelItemIds)

// 	if travelItemRepoErr != nil {
// 		panic(travelItemRepoErr)
// 	}
// }

// func (service *InvoiceDomainService) RemoveTravelItem(invoice Invoice, travelItems []*travelItemDomain.TravelItem) {

// 	imputationCount, err := service.imputationRepository.CountByInvoiceId(invoice.Id)

// 	if err != nil {
// 		panic(err)
// 	}

// 	if *imputationCount > 0 {
// 		panic(CustomErrors.DomainError(errors.New("cannot remove travel item to an invoice  with imputations")))
// 	}

// 	travelItemIds := make([]int, 0)

// 	for _, travelItem := range travelItems {
// 		travelItemIds = append(travelItemIds, int(travelItem.Id))
// 		invoice.Amount -= travelItem.TotalPrice
// 	}

// 	invoiceRepoErr := service.invoiceRepository.Update(service.transactionManager.GetTransaction(), &invoice)

// 	if invoiceRepoErr != nil {
// 		panic(invoiceRepoErr)
// 	}

// 	travelItemRepoErr := service.travelItemRepository.UpdateByInvoiceId(service.transactionManager.GetTransaction(), nil, travelItemIds)

// 	if travelItemRepoErr != nil {
// 		panic(travelItemRepoErr)
// 	}
// }

// func (service *InvoiceDomainService) VoidInvoice(invoiceId types.EID) error {
// 	defer service.transactionManager.CatchError()

// 	return nil
// }
