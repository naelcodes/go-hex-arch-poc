package application

import (
	"fmt"

	paymentDomain "github.com/naelcodes/ab-backend/internal/core/domains/payment-domain"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	"github.com/naelcodes/ab-backend/pkg/types"
	"github.com/naelcodes/ab-backend/pkg/utils"
)

func (application *Application) GetAllPaymentsService(queryParams *types.GetQueryParams) (*dto.GetAllPaymentsDTO, error) {

	totalRowCount, err := application.paymentRepository.Count()

	if err != nil {
		return nil, err
	}

	if queryParams == nil || (queryParams.PageNumber == nil && queryParams.PageSize == nil) {
		if queryParams == nil {
			queryParams = new(types.GetQueryParams)
		}
		queryParams.PageNumber = new(int)
		queryParams.PageSize = new(int)
		*queryParams.PageNumber = 0
		*queryParams.PageSize = *totalRowCount
	}

	payments, err := application.paymentRepository.GetAll(queryParams)

	if err != nil {
		return nil, err
	}

	getPaymentsDTO := new(dto.GetAllPaymentsDTO)
	getPaymentsDTO.Data = payments
	getPaymentsDTO.PageNumber = *queryParams.PageNumber
	getPaymentsDTO.PageSize = *queryParams.PageSize
	getPaymentsDTO.TotalRowCount = *totalRowCount

	return getPaymentsDTO, nil

}

func (application *Application) GetPaymentService(id types.EID) (*dto.GetPaymentDTO, error) {

	payment, err := application.paymentRepository.GetById(types.EID(id))
	if err != nil {
		return nil, err
	}

	return payment, nil
}

func (application *Application) CreatePaymentService(paymentDTO *dto.CreatePaymentDTO) (*dto.GetPaymentDTO, error) {

	totalRowCount, err := application.paymentRepository.Count()
	if err != nil {
		utils.Logger.Error(fmt.Sprintf("payment records Count error: %v", err))
		return nil, err
	}

	utils.Logger.Info(fmt.Sprintf("payment records Count: %v", totalRowCount))

	payment := paymentDomain.NewPaymentBuilder().
		SetAmount(paymentDTO.Amount).
		SetIdCustomer(types.EID(paymentDTO.IdCustomer)).
		SetPaymentMode(paymentDTO.PaymentMode).
		SetPaymentNumber(*totalRowCount + 1).
		SetPaymentDate().
		SetBalance(paymentDTO.Amount - 10).
		SetUsedAmount(10).
		Build()

	savedPaymentDTO, err := application.paymentRepository.Save(payment)

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("payment save error: %v", err))
		return nil, err
	}

	utils.Logger.Info(fmt.Sprintf("payment saved: %v", savedPaymentDTO))

	return savedPaymentDTO, nil

}
