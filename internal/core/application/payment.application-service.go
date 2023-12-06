package application

import (
	"fmt"

	paymentDomain "github.com/naelcodes/ab-backend/internal/core/domains/payment-domain"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	"github.com/naelcodes/ab-backend/pkg/types"
	"github.com/naelcodes/ab-backend/pkg/utils"
)

func (application *Application) GetAllPaymentsService(queryParams *types.GetQueryParams) (*dto.GetAllPaymentsDTO, error) {

	utils.Logger.Info(fmt.Sprintf("[GetAllPaymentsService] Query params: %v", queryParams))

	getAllPaymentsDTO, err := application.paymentRepository.GetAll(queryParams)

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[GetAllPaymentsService] Error getting all payments: %v", err))
		return nil, err
	}

	utils.Logger.Info(fmt.Sprintf("[GetAllPaymentsService] Get all payments DTO: %v", getAllPaymentsDTO))

	return getAllPaymentsDTO, nil

}

func (application *Application) GetPaymentService(id types.EID) (*dto.GetPaymentDTO, error) {

	utils.Logger.Info(fmt.Sprintf("[GetPaymentService] Payment ID: %v", id))

	paymentDTO, err := application.paymentRepository.GetById(types.EID(id))
	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[GetPaymentService] Error getting payment: %v", err))
		return nil, err
	}

	utils.Logger.Info(fmt.Sprintf("[GetPaymentService] Payment DTO: %v", paymentDTO))
	return paymentDTO, nil
}

func (application *Application) CreatePaymentService(paymentDTO *dto.CreatePaymentDTO) (*dto.GetPaymentDTO, error) {

	utils.Logger.Info(fmt.Sprintf("payment DTO: %v", paymentDTO))

	paymentBuilder := paymentDomain.NewPaymentBuilder().
		SetAmount(paymentDTO.Amount).
		SetIdCustomer(types.EID(paymentDTO.IdCustomer)).
		SetPaymentMode(paymentDTO.PaymentMode).
		SetPaymentDate().
		SetBalance(paymentDTO.Amount).
		SetUsedAmount(0)

	err := paymentBuilder.Validate()

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("payment validation error: %v", err))
		return nil, err
	}

	paymentDomainModel := paymentBuilder.Build()

	utils.Logger.Info(fmt.Sprintf("payment domain model: %v", paymentDomainModel))

	savedPaymentDTO, err := application.paymentRepository.Save(paymentDomainModel)

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("payment saving error: %v", err))
		return nil, err
	}

	utils.Logger.Info(fmt.Sprintf("saved payment DTO: %v", savedPaymentDTO))
	return savedPaymentDTO, nil

}
