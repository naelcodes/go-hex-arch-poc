package application

import (
	"errors"
	"fmt"

	invoiceDomain "github.com/naelcodes/ab-backend/internal/core/domains/invoice-domain"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	CustomError "github.com/naelcodes/ab-backend/pkg/errors"
	"github.com/naelcodes/ab-backend/pkg/types"
	"github.com/naelcodes/ab-backend/pkg/utils"
)

func (application *Application) CreateInvoiceService(createInvoiceDto *dto.CreateInvoiceDTO) (*dto.GetInvoiceDTO, error) {

	utils.Logger.Info(fmt.Sprintf("[CreateInvoiceService] - CreateInvoiceDTO: %v", createInvoiceDto))

	travelItemsId := make([]int, 0)
	invoiceAmount := float64(0)

	for _, travelItem := range createInvoiceDto.TravelItems {
		travelItemsId = append(travelItemsId, travelItem.Id)
		invoiceAmount += travelItem.TotalPrice
	}

	invoiceBuilder := invoiceDomain.NewInvoiceBuilder().
		SetCreationDate(createInvoiceDto.CreationDate).
		SetDueDate(createInvoiceDto.DueDate).
		SetIdCustomer(types.EID(createInvoiceDto.IdCustomer)).
		SetAmount(invoiceAmount).
		SetCreditApply(0).
		SetBalance(invoiceAmount).
		SetTravelItemsId(travelItemsId)

	err := invoiceBuilder.Validate()

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[CreateInvoiceService] - Error validating invoice: %v", err))
		return nil, err
	}

	invoice := invoiceBuilder.Build()

	transactionErr := application.TransactionManager.Begin()

	if transactionErr != nil {
		utils.Logger.Error(fmt.Sprintf("[CreateInvoiceService] - Error starting transaction: %v", transactionErr))
		return nil, CustomError.ServiceError(transactionErr, "TransactionManager.Begin()")
	}

	savedInvoiceDto, repoError := application.invoiceRepository.Save(application.TransactionManager.GetTransaction(), invoice)

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
