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
	"github.com/naelcodes/ab-backend/pkg/utils"
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

	utils.Logger.Info(fmt.Sprintf("[PaymentRepository - CountByCustomerID] Customer ID: %v", customerId))

	totalRowCount, err := repo.Database.Payment.Query().
		Where(payment.
			And(
				payment.HasCustomerWith(customer.IDEQ(int(customerId))),
				payment.TagEQ(payment.Tag3))).
		Count(repo.Context)

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[PaymentRepository - CountByCustomerID] Error counting customer's payments: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error counting customer's payments: %v", err))
	}

	utils.Logger.Info(fmt.Sprintf("[PaymentRepository - CountByCustomerID] Total number of customer's payments: %v", totalRowCount))
	return &totalRowCount, nil
}

func (repo *PaymentRepository) GetById(id types.EID) (*dto.GetPaymentDTO, error) {

	utils.Logger.Info(fmt.Sprintf("[PaymentRepository - GetById] Payment ID: %v", id))

	payment, err := repo.Database.Payment.Query().Where(payment.IDEQ(int(id))).
		WithCustomer(func(q *ent.CustomerQuery) {
			q.Select(customer.FieldID)
		}).First(repo.Context)

	if err != nil {
		if ent.IsNotFound(err) {
			utils.Logger.Error("[PaymentRepository - GetById] Payment not found")
			return nil, CustomErrors.RepositoryError(errors.New("payment record not found"))
		}

		utils.Logger.Error(fmt.Sprintf("[PaymentRepository - GetById] Error getting payment: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error getting payment record: %v", err))
	}

	utils.Logger.Info(fmt.Sprintf("[PaymentRepository - GetById] Payment: %v", payment))

	paymentDTO := PaymentModelToDTO(payment, false, nil)

	utils.Logger.Info(fmt.Sprintf("[PaymentRepository - GetById] Payment DTO: %v", paymentDTO))
	return paymentDTO, nil
}

func (repo *PaymentRepository) GetAll(queryParams *types.GetQueryParams) (*dto.GetAllPaymentsDTO, error) {
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

	totalRowCount, err := repo.Count()
	if err != nil {
		return nil, err
	}

	pageNumber := 0
	pageSize := *totalRowCount

	if queryParams != nil && queryParams.PageNumber != nil && queryParams.PageSize != nil {
		pageNumber = *queryParams.PageNumber
		pageSize = *queryParams.PageSize
		PaymentQuery.Offset(pageNumber * pageSize).Limit(pageSize)
	}

	payments, err := PaymentQuery.All(repo.Context)

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[PaymentRepository - GetAll] Error getting payments: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error getting payments records: %v", err))
	}

	utils.Logger.Info(fmt.Sprintf("[PaymentRepository - GetAll] Total number of payments: %v", totalRowCount))
	paymentDTOList := PaymentModelListToDTOList(payments, embedCustomer, nil)

	utils.Logger.Info(fmt.Sprintf("[PaymentRepository - GetAll] Payments DTO: %v", paymentDTOList))

	getAllPaymentsDTO := new(dto.GetAllPaymentsDTO)
	getAllPaymentsDTO.Data = paymentDTOList
	getAllPaymentsDTO.TotalRowCount = *totalRowCount
	getAllPaymentsDTO.PageNumber = pageNumber
	getAllPaymentsDTO.PageSize = pageSize

	utils.Logger.Info(fmt.Sprintf("[PaymentRepository - GetAll] GetAllPaymentsDTO: %v", getAllPaymentsDTO))

	return getAllPaymentsDTO, nil
}

func (repo *PaymentRepository) Save(paymentDomainModel *paymentDomain.Payment) (*dto.GetPaymentDTO, error) {

	utils.Logger.Info(fmt.Sprintf("Reposiotry - Saving payment domain model: %v", paymentDomainModel))

	totalRowCount, err := repo.Count()
	if err != nil {
		return nil, err
	}

	payment, err := repo.Database.Payment.Create().
		SetAmount(paymentDomainModel.Amount).
		SetCustomerID(int(paymentDomainModel.IdCustomer)).
		SetBalance(paymentDomainModel.Balance).
		SetUsedAmount(paymentDomainModel.UsedAmount).
		SetFop(payment.Fop(paymentDomainModel.PaymentMode)).
		SetDate(paymentDomainModel.PaymentDate).
		SetNumber(utils.GenerateCode("pr", *totalRowCount+1)).
		Save(repo.Context)

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("Repository - Error saving payment: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error saving payment: %v", err))
	}

	utils.Logger.Info(fmt.Sprintf("Repository - Converting to DTO: %v", payment))

	customerId := int(paymentDomainModel.IdCustomer)
	paymentDTO := PaymentModelToDTO(payment, false, &customerId)

	utils.Logger.Info(fmt.Sprintf("Repository - Saved payment DTO: %v", paymentDTO))

	return paymentDTO, nil

}

func (repo *PaymentRepository) SaveAllPaymentsAllocations(transaction *ent.Tx, payments []*paymentDomain.Payment) {

	utils.Logger.Info(fmt.Sprintf("[PaymentRepository - SavePaymentsAllocations] - Saving payments allocations: %v", payments))

	for _, p := range payments {
		updatedPayment, err := transaction.Payment.UpdateOneID(int(p.Id)).
			SetBalance(p.Balance).
			SetUsedAmount(p.UsedAmount).
			SetStatus(payment.Status(p.Status)).
			Save(repo.Context)

		utils.Logger.Info(fmt.Sprintf("[PaymentRepository - SavePaymentsAllocations] - Updated payment: %v", updatedPayment))
		if err != nil {
			panic(CustomErrors.RepositoryError(fmt.Errorf("error saving payments allocations: %v", err)))
		}
	}

	utils.Logger.Info("[PaymentRepository - SavePaymentsAllocations] - Saved payments allocations")
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
