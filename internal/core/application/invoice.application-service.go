package application

import (
	"errors"
	"fmt"

	imputationDomain "github.com/naelcodes/ab-backend/internal/core/domains/imputation-domain"
	invoiceDomain "github.com/naelcodes/ab-backend/internal/core/domains/invoice-domain"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	CustomError "github.com/naelcodes/ab-backend/pkg/errors"
	"github.com/naelcodes/ab-backend/pkg/types"
	"github.com/naelcodes/ab-backend/pkg/utils"
)

func (application *Application) CreateInvoiceService(createInvoiceDto *dto.CreateInvoiceDTO) (*dto.GetInvoiceDTO, error) {

	utils.Logger.Info(fmt.Sprintf("[CreateInvoiceService] - CreateInvoiceDTO: %v", createInvoiceDto))

	invoiceAmount := float64(0)
	travelItemIdList := make([]int, 0)

	//check if travelItems exist and it's not used
	for _, travelItem := range createInvoiceDto.TravelItemIds {
		travelItemDTO, err := application.travelItemRepository.GetById(types.EID(travelItem.Id))

		if err != nil {
			utils.Logger.Error(fmt.Sprintf("[CreateInvoiceService] - Error getting travel item: %v", err))
			return nil, err
		}

		if travelItemDTO.IdInvoice != nil {
			utils.Logger.Error(fmt.Sprintf("[CreateInvoiceService] - Travel item already has an invoice: %v", travelItemDTO))
			return nil, CustomError.ValidationError(errors.New("travel item already has an invoice"))
		}

		travelItemIdList = append(travelItemIdList, int(travelItem.Id))
		invoiceAmount += *travelItemDTO.TotalPrice
	}

	invoiceBuilder := invoiceDomain.NewInvoiceBuilder().
		SetCreationDate(createInvoiceDto.CreationDate).
		SetDueDate(createInvoiceDto.DueDate).
		SetIdCustomer(types.EID(createInvoiceDto.IdCustomer)).
		SetAmount(invoiceAmount).
		SetCreditApply(0).
		SetBalance(invoiceAmount).
		SetTravelItemsId(travelItemIdList)

	err := invoiceBuilder.Validate()

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[CreateInvoiceService] - Error validating invoice: %v", err))
		return nil, err
	}

	invoiceDomainModel := invoiceBuilder.Build()

	transactionErr := application.TransactionManager.Begin()

	if transactionErr != nil {
		utils.Logger.Error(fmt.Sprintf("[CreateInvoiceService] - Error starting transaction: %v", transactionErr))
		return nil, CustomError.ServiceError(transactionErr, "TransactionManager.Begin()")
	}

	savedInvoiceDto, repoError := application.invoiceRepository.Save(application.TransactionManager.GetTransaction(), invoiceDomainModel)

	if repoError != nil {
		utils.Logger.Error(fmt.Sprintf("[CreateInvoiceService] - Error saving invoice: %v", repoError))
		RollbackErr := application.TransactionManager.Rollback()
		if RollbackErr != nil {
			utils.Logger.Error(fmt.Sprintf("[CreateInvoiceService] - Error rolling back transaction: %v", RollbackErr))
			return nil, CustomError.ServiceError(errors.Join(RollbackErr, repoError), "TransactionManager.Rollback()")
		}
		return nil, repoError
	}

	transactionCommitErr := application.TransactionManager.Commit()

	if transactionErr != nil {
		utils.Logger.Error(fmt.Sprintf("[CreateInvoiceService] - Error commiting transaction: %v", transactionCommitErr))
		return nil, CustomError.ServiceError(transactionCommitErr, "TransactionManager.Commit()")
	}

	utils.Logger.Info("[CreateInvoiceService] - Transaction committed")
	utils.Logger.Info(fmt.Sprintf("[CreateInvoiceService] - Invoice created: %v", savedInvoiceDto))

	return savedInvoiceDto, nil
}

// Get all invoices
func (application *Application) GetAllInvoiceService(queryParams *types.GetQueryParams) (*dto.GetAllInvoiceDTO, error) {

	utils.Logger.Info(fmt.Sprintf("[GetAllInvoicesService] - GetQueryParams: %v", queryParams))

	getAllInvoiceDTO, err := application.invoiceRepository.GetAll(queryParams)

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[GetAllInvoicesService] - Error getting invoices: %v", err))
		return nil, err
	}

	utils.Logger.Info(fmt.Sprintf("[GetAllInvoicesService] - GetAllInvoiceDTO: %v", getAllInvoiceDTO))
	return getAllInvoiceDTO, nil

}

// Get invoice
func (application *Application) GetInvoiceService(id int, queryParams *types.GetQueryParams) (*dto.GetInvoiceDTO, error) {

	utils.Logger.Info(fmt.Sprintf("[GetInvoiceService] - GetInvoiceDTO: %v", id))

	getInvoiceDTO, err := application.invoiceRepository.GetById(types.EID(id), queryParams)

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[GetInvoiceService] - Error getting invoice: %v", err))
		return nil, err
	}

	utils.Logger.Info(fmt.Sprintf("[GetInvoiceService] - GetInvoiceDTO: %v", getInvoiceDTO))

	return getInvoiceDTO, nil

}

func (application *Application) ApplyInvoiceImputationService(invoiceId types.EID, invoiceImputationDTOList []*dto.InvoiceImputationDTO) (*dto.ImputationOperationResult, error) {

	utils.Logger.Info(fmt.Sprintf("[ApplyInvoiceImputationService] - InvoiceImputationDTOList: %v", invoiceImputationDTOList))
	imputationOperationResult := new(dto.ImputationOperationResult)

	exists, err := application.invoiceRepository.Exists(invoiceId)

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[ApplyInvoiceImputationService] - Error checking if invoice exists: %v", err))
		return nil, err
	}

	if !exists {
		utils.Logger.Error(fmt.Sprintf("[ApplyInvoiceImputationService] - Invoice does not exist: %v", invoiceId))
		return nil, CustomError.ValidationError(errors.New("invoice does not exist"))
	}

	//construct imputationDomainModel

	ImputationDomainModelList := make([]*imputationDomain.Imputation, 0)
	PaymentsIdList := make([]int, 0)

	for _, invoiceImputationDTO := range invoiceImputationDTOList {
		ImputationDomainModelList = append(ImputationDomainModelList, imputationDomain.NewImputationBuilder().
			SetIdInvoice(invoiceId).
			SetIdPayment(types.EID(invoiceImputationDTO.IdPayment)).
			SetAmountApplied(invoiceImputationDTO.AmountApplied).
			Build())

		PaymentsIdList = append(PaymentsIdList, invoiceImputationDTO.IdPayment)
	}

	notFoundList, err := application.paymentRepository.CheckInvoiceOwnerPayments(invoiceId, PaymentsIdList)

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[ApplyInvoiceImputationService] - Error checking if payment exists: %v", err))
		return nil, err
	}

	if len(*notFoundList) > 0 {
		utils.Logger.Error(fmt.Sprintf("[ApplyInvoiceImputationService] - %v Payments used in imputation does not belong to invoice owner, payment ids: %v", len(*notFoundList), *notFoundList))
		return nil, CustomError.ValidationError(fmt.Errorf("%v payments used in imputation does not belong to invoice owner, payment ids: %v", len(*notFoundList), *notFoundList))
	}

	transactionErr := application.TransactionManager.Begin()

	if transactionErr != nil {
		utils.Logger.Error(fmt.Sprintf("[ApplyInvoiceImputationService] - Error starting transaction: %v", transactionErr))
		return nil, CustomError.ServiceError(transactionErr, "TransactionManager.Begin()")
	}

	domainService := invoiceDomain.NewInvoiceDomainService(application.imputationRepository, application.paymentRepository, application.invoiceRepository, application.TransactionManager)

	insertedCount, updatedCount, deletedCount, err := domainService.ApplyImputation(invoiceId, ImputationDomainModelList)

	imputationOperationResult.InsertedImputationCount = insertedCount
	imputationOperationResult.UpdatedImputationCount = updatedCount
	imputationOperationResult.DeletedImputationCount = deletedCount

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[ApplyInvoiceImputationService] - Error applying imputation: %v", err))
		RollbackErr := application.TransactionManager.Rollback()

		if RollbackErr != nil {
			utils.Logger.Error(fmt.Sprintf("[ApplyInvoiceImputationService] - Error rolling back transaction: %v", RollbackErr))
			return nil, CustomError.ServiceError(errors.Join(RollbackErr, err), "TransactionManager.Rollback()")
		}
		return nil, err
	}

	transactionCommitErr := application.TransactionManager.Commit()

	if transactionErr != nil {
		utils.Logger.Error(fmt.Sprintf("[ApplyInvoiceImputationService] - Error committing transaction: %v", transactionCommitErr))
		return nil, CustomError.ServiceError(transactionCommitErr, "TransactionManager.Commit()")
	}

	utils.Logger.Info("[ApplyInvoiceImputationService] - Transaction committed")

	return imputationOperationResult, nil

}

func (application *Application) GetInvoiceImputationService(invoiceId types.EID) (*dto.GetInvoiceImputationDTO, error) {

	utils.Logger.Info(fmt.Sprintf("[GetInvoiceImputationService] - InvoiceId: %v", invoiceId))

	getInvoiceImputationDTO, err := application.imputationRepository.GetByInvoiceId(invoiceId)

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[GetInvoiceImputationService] - Error getting invoice imputation: %v", err))
		return nil, err
	}

	utils.Logger.Info(fmt.Sprintf("[GetInvoiceImputationService] - GetInvoiceImputationDTO: %v", getInvoiceImputationDTO))

	return getInvoiceImputationDTO, nil

}
