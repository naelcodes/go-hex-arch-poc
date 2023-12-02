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
	"github.com/naelcodes/ab-backend/pkg/types"
)

type PaymentRepository struct {
	Database *ent.Client
	Context  context.Context
}

func (repo *PaymentRepository) Count() (*int, error) {

	totalRowCount, err := repo.Database.Payment.Query().Where(payment.TagEQ(payment.Tag3)).Count(repo.Context)

	if err != nil {
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error counting payments: %v", err))
	}
	return &totalRowCount, nil
}

func (repo *PaymentRepository) CountByCustomerID(customerId types.EID) (*int, error) {

	totalRowCount, err := repo.Database.Payment.Query().
		Where(payment.
			And(
				payment.HasCustomerWith(customer.IDEQ(int(customerId))),
				payment.TagEQ(payment.Tag3))).
		Count(repo.Context)

	if err != nil {
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error counting customer's payments: %v", err))
	}
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
	paymentDTO := PaymentModelToDTO(payment, false)
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

	paymentDTOList := PaymentModelListToDTOList(payments, embedCustomer)

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

	paymentDTOList := PaymentModelListToDTOList(payments, false)

	return paymentDTOList, nil

}

func (repo *PaymentRepository) Save(paymentAggregate *paymentDomain.PaymentAggregate) (*dto.GetPaymentDTO, error) {

	payment, err := repo.Database.Payment.Create().
		SetAmount(paymentAggregate.Amount).
		SetCustomerID(int(paymentAggregate.IdCustomer)).
		SetBalance(paymentAggregate.Balance).
		SetUsedAmount(paymentAggregate.UsedAmount).
		SetFop(payment.Fop(paymentAggregate.PaymentMode)).
		SetStatus(payment.Status(paymentAggregate.Status)).
		SetDate(paymentAggregate.PaymentDate).
		SetNumber(paymentAggregate.PaymentNumber).
		Save(repo.Context)

	if err != nil {
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error saving payment: %v", err))
	}

	paymentDTO := PaymentModelToDTO(payment, false)
	return paymentDTO, nil

}

func (repo *PaymentRepository) Update(tx *ent.Tx, paymentAggregate *paymentDomain.PaymentAggregate) error {
	if tx != nil {
		_, err := tx.Payment.UpdateOneID(int(paymentAggregate.Id)).
			SetAmount(paymentAggregate.Amount).
			SetBalance(paymentAggregate.Balance).
			SetUsedAmount(paymentAggregate.UsedAmount).
			SetStatus(payment.Status(paymentAggregate.Status)).
			Save(repo.Context)

		if err != nil {
			return CustomErrors.RepositoryError(fmt.Errorf("error updating payment: %v", err))
		}
	} else {
		paymentQuery := repo.Database.Payment.UpdateOneID(int(paymentAggregate.Id))
		if paymentAggregate.IdCustomer != 0 {
			paymentQuery.SetCustomerID(int(paymentAggregate.IdCustomer))
		}

		if paymentAggregate.Amount != 0 {
			paymentQuery.SetAmount(paymentAggregate.Amount)
		}

		if paymentAggregate.PaymentMode != "" {
			paymentQuery.SetFop(payment.Fop(paymentAggregate.PaymentMode))
		}

		_, err := paymentQuery.Save(repo.Context)

		if err != nil {
			return CustomErrors.RepositoryError(fmt.Errorf("error updating payment: %v", err))
		}
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
