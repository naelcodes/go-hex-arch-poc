package postgres

import (
	"context"
	"errors"
	"fmt"

	"github.com/naelcodes/ab-backend/ent"
	"github.com/naelcodes/ab-backend/ent/customer"
	"github.com/naelcodes/ab-backend/ent/payment"
	paymentDomain "github.com/naelcodes/ab-backend/internal/core/domains/payment-domain"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	CustomErrors "github.com/naelcodes/ab-backend/pkg/errors"
	"github.com/naelcodes/ab-backend/pkg/logger"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type PaymentRepository struct {
	Database *ent.Client
	Context  context.Context
	Logger   *logger.Logger
}

func (repo *PaymentRepository) Count() (*int, error) {

	totalRowCount, err := repo.Database.Payment.Query().Where(payment.TagEQ(payment.Tag3)).Count(repo.Context)

	if err != nil {
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error counting payments: %v", err))
	}
	return &totalRowCount, nil
}

func (repo *PaymentRepository) CountByCustomerID(customerId types.EID) (*int, error) {

	repo.Logger.Info(fmt.Sprintf("[PaymentRepository - CountByCustomerID] Customer ID: %v", customerId))

	totalRowCount, err := repo.Database.Payment.Query().
		Where(payment.
			And(
				payment.HasCustomerWith(customer.IDEQ(int(customerId))),
				payment.TagEQ(payment.Tag3))).
		Count(repo.Context)

	if err != nil {
		repo.Logger.Error(fmt.Sprintf("[PaymentRepository - CountByCustomerID] Error counting customer's payments: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error counting customer's payments: %v", err))
	}

	repo.Logger.Info(fmt.Sprintf("[PaymentRepository - CountByCustomerID] Total number of customer's payments: %v", totalRowCount))
	return &totalRowCount, nil
}

func (repo *PaymentRepository) GetById(id types.EID) (*dto.GetPaymentDTO, error) {

	payment, err := repo.Database.Payment.Query().Where(payment.IDEQ(int(id))).
		WithCustomer(func(q *ent.CustomerQuery) {
			q.Select(customer.FieldID)
		}).First(repo.Context)

	if err != nil {
		if ent.IsNotFound(err) {
			return nil, CustomErrors.RepositoryError(errors.New("payment record not found"))
		}
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error getting payment record: %v", err))
	}
	paymentDTO := PaymentModelToDTO(payment, false, nil)
	return paymentDTO, nil
}

func (repo *PaymentRepository) GetAll(queryParams *types.GetQueryParams) ([]*dto.GetPaymentDTO, error) {
	embedCustomer := false
	PaymentQuery := repo.Database.Payment.Query().Where(payment.TagEQ(payment.Tag3))

	if queryParams != nil && queryParams.Embed != nil && *queryParams.Embed == "customer" {
		PaymentQuery.WithCustomer()
		embedCustomer = true
	} else {
		PaymentQuery.WithCustomer(func(q *ent.CustomerQuery) {
			q.Select(customer.FieldID)
		})
	}

	if queryParams != nil && queryParams.PageNumber != nil && queryParams.PageSize != nil {
		PaymentQuery.Offset(*queryParams.PageNumber * *queryParams.PageSize).Limit(*queryParams.PageSize)
	}

	payments, err := PaymentQuery.All(repo.Context)

	if err != nil {
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error getting payments records: %v", err))
	}

	paymentDTOList := PaymentModelListToDTOList(payments, embedCustomer, nil)

	return paymentDTOList, nil
}

func (repo *PaymentRepository) GetByCustomerID(id types.EID, queryParams *types.GetQueryParams, isOpen *bool) ([]*dto.GetPaymentDTO, error) {

	PaymentQuery := repo.Database.Payment.Query()

	if isOpen != nil && *isOpen {
		PaymentQuery.Where(payment.
			And(payment.StatusEQ(payment.StatusOpen),
				payment.HasCustomerWith(customer.IDEQ(int(id))),
				payment.TagEQ(payment.Tag3)))
	} else {
		PaymentQuery.Where(payment.
			And(payment.HasCustomerWith(customer.IDEQ(int(id))),
				payment.TagEQ(payment.Tag3)))
	}

	if queryParams != nil && queryParams.PageNumber != nil && queryParams.PageSize != nil {
		PaymentQuery.Offset(*queryParams.PageNumber * *queryParams.PageSize).Limit(*queryParams.PageSize)
	}

	payments, err := PaymentQuery.All(repo.Context)

	if err != nil {
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error getting customer's payments records: %v", err))
	}

	customerId := int(id)
	paymentDTOList := PaymentModelListToDTOList(payments, false, &customerId)

	return paymentDTOList, nil

}

func (repo *PaymentRepository) Save(paymentEntity *paymentDomain.Payment) (*dto.GetPaymentDTO, error) {

	repo.Logger.Info(fmt.Sprintf("Reposiotry - Saving payment entity: %v", paymentEntity))

	payment, err := repo.Database.Payment.Create().
		SetAmount(paymentEntity.Amount).
		SetCustomerID(int(paymentEntity.IdCustomer)).
		SetBalance(paymentEntity.Balance).
		SetUsedAmount(paymentEntity.UsedAmount).
		SetFop(payment.Fop(paymentEntity.PaymentMode)).
		SetDate(paymentEntity.PaymentDate).
		SetNumber(paymentEntity.PaymentNumber).
		Save(repo.Context)

	if err != nil {
		repo.Logger.Error(fmt.Sprintf("Repository - Error saving payment: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error saving payment: %v", err))
	}

	repo.Logger.Info(fmt.Sprintf("Repository - Converting to DTO: %v", payment))

	customerId := int(paymentEntity.IdCustomer)
	paymentDTO := PaymentModelToDTO(payment, false, &customerId)

	repo.Logger.Info(fmt.Sprintf("Repository - Saved payment DTO: %v", paymentDTO))

	return paymentDTO, nil

}

func (repo *PaymentRepository) SaveAllPaymentsAllocations(transaction *ent.Tx, payments []*paymentDomain.Payment) {

	repo.Logger.Info(fmt.Sprintf("[PaymentRepository - SavePaymentsAllocations] - Saving payments allocations: %v", payments))

	for _, p := range payments {
		updatedPayment, err := transaction.Payment.UpdateOneID(int(p.Id)).
			SetBalance(p.Balance).
			SetUsedAmount(p.UsedAmount).
			SetStatus(payment.Status(p.Status)).
			Save(repo.Context)

		repo.Logger.Info(fmt.Sprintf("[PaymentRepository - SavePaymentsAllocations] - Updated payment: %v", updatedPayment))
		if err != nil {
			panic(CustomErrors.RepositoryError(fmt.Errorf("error saving payments allocations: %v", err)))
		}
	}

	repo.Logger.Info("[PaymentRepository - SavePaymentsAllocations] - Saved payments allocations")
}

func (repo *PaymentRepository) Update(paymentEntity *paymentDomain.Payment) error {

	paymentData, err := repo.Database.Payment.Query().
		WithCustomer(func(q *ent.CustomerQuery) {
			q.Select(customer.FieldID)
		}).Where(payment.IDEQ(int(paymentEntity.Id))).Only(repo.Context)

	if err != nil {
		return CustomErrors.RepositoryError(fmt.Errorf("error getting payment during update: %v", err))
	}

	paymentUpdateQuery := repo.Database.Payment.UpdateOneID(int(paymentEntity.Id))

	if paymentEntity.IdCustomer != types.EID(paymentData.Edges.Customer.ID) {
		paymentUpdateQuery.SetCustomerID(int(paymentEntity.IdCustomer))
	}

	if paymentEntity.Amount != paymentData.Amount {
		paymentUpdateQuery.SetAmount(paymentEntity.Amount)
	}

	if paymentEntity.PaymentMode != string(paymentData.Fop) {
		paymentUpdateQuery.SetFop(payment.Fop(paymentEntity.PaymentMode))
	}

	_, saveErr := paymentUpdateQuery.Save(repo.Context)

	if saveErr != nil {
		return CustomErrors.RepositoryError(fmt.Errorf("error updating payment: %v", err))
	}

	return nil

}

func (repo *PaymentRepository) Delete(id types.EID) error {
	deletedCount, err := repo.Database.Payment.Delete().Where(payment.IDEQ(int(id))).Exec(repo.Context)

	if err != nil {
		return CustomErrors.RepositoryError(fmt.Errorf("error deleting payment: %v", err))
	}

	if deletedCount == 0 {
		return CustomErrors.RepositoryError(errors.New("payment  record not found"))
	}

	return nil

}
