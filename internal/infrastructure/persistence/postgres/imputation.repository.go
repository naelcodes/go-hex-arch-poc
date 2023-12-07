package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/naelcodes/ab-backend/ent"
	"github.com/naelcodes/ab-backend/ent/customer"
	"github.com/naelcodes/ab-backend/ent/imputation"
	"github.com/naelcodes/ab-backend/ent/invoice"
	"github.com/naelcodes/ab-backend/ent/payment"
	imputationDomain "github.com/naelcodes/ab-backend/internal/core/domains/imputation-domain"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	CustomErrors "github.com/naelcodes/ab-backend/pkg/errors"
	"github.com/naelcodes/ab-backend/pkg/types"
	"github.com/naelcodes/ab-backend/pkg/utils"
)

type ImputationRepository struct {
	Database *ent.Client
	Context  context.Context
}

func (repo *ImputationRepository) CountByInvoiceId(idInvoice types.EID) (*int, error) {
	return nil, nil
}

func (repo *ImputationRepository) CountByPaymentId(idPayment types.EID) (*int, error) {
	return nil, nil
}

func (repo *ImputationRepository) GetByInvoiceId(idInvoice types.EID) (*dto.GetInvoiceImputationDTO, error) {

	utils.Logger.Info(fmt.Sprintf("[ImputationRepository - GetByInvoiceId] Invoice ID: %v", idInvoice))

	invoiceRecord, err := repo.Database.Invoice.Query().WithCustomer(func(q *ent.CustomerQuery) {
		q.Select(customer.FieldID)
	}).Where(invoice.IDEQ(int(idInvoice))).Only(repo.Context)

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[ImputationRepository - GetByInvoiceId] Error getting invoice: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error getting invoice: %v", err))
	}

	customerId := int(invoiceRecord.Edges.Customer.ID)
	utils.Logger.Info(fmt.Sprintf("[ImputationRepository - GetByInvoiceId] Customer ID: %v", customerId))

	paymentRecords, err := repo.Database.Payment.Query().
		Where(payment.
			And(payment.HasCustomerWith(customer.IDEQ(customerId)),
				// payment.StatusEQ(payment.StatusOpen),
				payment.TagEQ(payment.Tag3))).
		All(repo.Context)

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[ImputationRepository - GetByInvoiceId] Error getting payments: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error getting payments: %v", err))
	}

	imputationDetailsList := make([]*dto.GetImputationDetails, 0)

	for _, paymentRecord := range paymentRecords {

		imputationDetails := new(dto.GetImputationDetails)

		paymentDetails := new(dto.PaymentDetails)
		paymentDetails.Id = int(paymentRecord.ID)
		paymentDetails.Amount = paymentRecord.Amount
		paymentDetails.Balance = paymentRecord.Balance
		paymentDetails.PaymentDate = utils.FormatDate(paymentRecord.Date)
		paymentDetails.PaymentNumber = paymentRecord.Number

		imputationDetails.Payment = paymentDetails
		imputationDetails.AmountApplied = 0

		imputationRecord, err := repo.Database.Imputation.Query().
			Where(imputation.
				And(imputation.HasPaymentWith(payment.IDEQ(int(paymentRecord.ID))),
					imputation.HasInvoiceWith(invoice.IDEQ(int(idInvoice))),
					imputation.TagEQ(imputation.Tag3))).
			Only(repo.Context)

		if err != nil && !ent.IsNotFound(err) {
			utils.Logger.Error(fmt.Sprintf("[ImputationRepository - GetByInvoiceId] Error getting imputation: %v", err))
			return nil, CustomErrors.RepositoryError(fmt.Errorf("error getting imputation: %v", err))
		}

		if imputationRecord != nil {
			imputationDetails.AmountApplied = imputationRecord.AmountApply
		}

		utils.Logger.Info(fmt.Sprintf("[ImputationRepository - GetByInvoiceId] Payment ID: %v, Imputation: %v", paymentRecord.ID, imputationDetails))
		imputationDetailsList = append(imputationDetailsList, imputationDetails)

	}

	getInvoiceImputationDTO := new(dto.GetInvoiceImputationDTO)
	getInvoiceImputationDTO.InvoiceAmount = invoiceRecord.Amount
	getInvoiceImputationDTO.Imputations = imputationDetailsList

	utils.Logger.Info(fmt.Sprintf("[ImputationRepository - GetByInvoiceId] Invoice ID: %v, Imputations: %v", idInvoice, getInvoiceImputationDTO))

	return getInvoiceImputationDTO, nil

}

func (repo *ImputationRepository) GetByPaymentAndInvoiceId(idPayment types.EID, idInvoice types.EID) (*bool, *ent.Imputation, error) {

	utils.Logger.Info(fmt.Sprintf("[ImputationRepository - GetByPaymentAndInvoiceId] Payment ID: %v, Invoice ID: %v", idPayment, idInvoice))

	exists := false
	imputation, err := repo.Database.Imputation.Query().
		WithPayment(func(q *ent.PaymentQuery) {
			q.Select(payment.FieldID)
		}).Where(
		imputation.And(
			imputation.HasPaymentWith(payment.IDEQ(int(idPayment))),
			imputation.HasInvoiceWith(invoice.IDEQ(int(idInvoice))),
			imputation.TagEQ(imputation.Tag3),
		)).
		Only(repo.Context)

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[ImputationRepository - GetByPaymentAndInvoiceId] Error getting imputation: %v", err))
		if ent.IsNotFound(err) {
			return &exists, nil, nil
		}
		if ent.IsNotSingular(err) {
			return nil, nil, CustomErrors.RepositoryError(fmt.Errorf("error getting  single imputation: %v", err))
		}
		return nil, nil, CustomErrors.RepositoryError(fmt.Errorf("error getting imputation: %v", err))
	}

	exists = true

	utils.Logger.Info(fmt.Sprintf("[ImputationRepository - GetByPaymentAndInvoiceId] Imputation found: %v", imputation))
	return &exists, imputation, nil
}

func (repo *ImputationRepository) Update(transaction *ent.Tx, imputationDomainModel *imputationDomain.Imputation) (int, error) {

	utils.Logger.Info(fmt.Sprintf("[ImputationRepository - Update] Updating imputation: %v", imputationDomainModel))

	updatedImputation, err := transaction.Imputation.UpdateOneID(int(imputationDomainModel.Id)).
		SetAmountApply(imputationDomainModel.AmountApplied).
		Save(repo.Context)

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[ImputationRepository - Update] Error updating imputation: %v", err))
		return 0, CustomErrors.RepositoryError(fmt.Errorf("error updating imputation: %v", err))
	}

	utils.Logger.Info(fmt.Sprintf("[ImputationRepository - Update] Updated imputation: %v", updatedImputation))
	return 1, nil
}

func (repo *ImputationRepository) SaveAll(transaction *ent.Tx, imputationDomainModelList []*imputationDomain.Imputation) (int, error) {

	utils.Logger.Info(fmt.Sprintf("[ImputationRepository - SaveAll] Saving %v imputations", len(imputationDomainModelList)))
	insertedCount := 0

	savedImputations, err := transaction.Imputation.MapCreateBulk(imputationDomainModelList, func(c *ent.ImputationCreate, index int) {
		c.SetInvoiceID(int(imputationDomainModelList[index].IdInvoice))
		c.SetPaymentID(int(imputationDomainModelList[index].IdPayment))
		c.SetAmountApply(imputationDomainModelList[index].AmountApplied)
	}).Save(repo.Context)

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[ImputationRepository - SaveAll] Error saving imputations: %v", err))
		return insertedCount, CustomErrors.RepositoryError(fmt.Errorf("error saving imputations: %v", err))
	}

	insertedCount = len(savedImputations)
	utils.Logger.Info(fmt.Sprintf("[ImputationRepository - SaveAll] Saved %v imputations", len(savedImputations)))

	return insertedCount, nil
}

func (repo *ImputationRepository) Delete(transaction *ent.Tx, id types.EID) (int, error) {

	utils.Logger.Info(fmt.Sprintf("[ImputationRepository - Delete] Deleting imputation: %v", id))

	deletedCount, err := transaction.Imputation.Delete().Where(imputation.IDEQ(int(id))).Exec(repo.Context)

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[ImputationRepository - Delete] Error deleting imputation: %v", err))
		return 0, CustomErrors.RepositoryError(fmt.Errorf("error deleting imputation: %v", err))
	}

	if deletedCount == 0 {
		utils.Logger.Error(fmt.Sprintf("[ImputationRepository - Delete] Imputation not found: %v", id))
		return 0, CustomErrors.RepositoryError(errors.New("imputation not found"))
	}

	utils.Logger.Info(fmt.Sprintf("[ImputationRepository - Delete] Deleted imputation: %v", id))
	return deletedCount, nil
}
