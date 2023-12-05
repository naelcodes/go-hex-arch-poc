package application

import (
	"errors"
	"fmt"

	travelItemDomain "github.com/naelcodes/ab-backend/internal/core/domains/TravelItem-domain"
	invoiceDomain "github.com/naelcodes/ab-backend/internal/core/domains/invoice-domain"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	"github.com/naelcodes/ab-backend/pkg/types"
)

func (application *Application) CreateInvoiceService(createInvoiceDto *dto.CreateInvoiceDTO) (*dto.GetInvoiceDTO, error) {

	defer application.TransactionManager.CatchError()

	application.Logger.Info(fmt.Sprintf("[CreateInvoiceService] - CreateInvoiceDTO: %v", createInvoiceDto))

	totalInvoiceCount, err := application.invoiceRepository.Count()

	if err != nil {
		application.Logger.Error(fmt.Sprintf("[CreateInvoiceService] - Error counting invoices: %v", err))
		return nil, err
	}

	application.Logger.Info(fmt.Sprintf("[CreateInvoiceService] - Total number of invoices: %v", totalInvoiceCount))

	invoice := invoiceDomain.NewInvoiceBuilder().
		SetCreationDate(createInvoiceDto.CreationDate).
		SetInvoiceNumber(*totalInvoiceCount + 1).
		SetDueDate(createInvoiceDto.DueDate).
		SetIdCustomer(types.EID(createInvoiceDto.IdCustomer)).
		SetAmount(0).
		SetCreditApply(0).
		SetBalance(0).
		Build()

	transactionErr := application.TransactionManager.Begin()

	if transactionErr != nil {
		application.Logger.Error(fmt.Sprintf("[CreateInvoiceService] - Error starting transaction: %v", transactionErr))
		return nil, transactionErr
	}

	savedInvoiceDto, saveInvoiceErr := application.invoiceRepository.Save(application.TransactionManager.GetTransaction(), invoice)

	if saveInvoiceErr != nil {
		err := application.TransactionManager.Rollback()
		application.Logger.Error(fmt.Sprintf("[CreateInvoiceService] - Error saving invoice: %v", errors.Join(err, saveInvoiceErr)))
		return nil, errors.Join(err, saveInvoiceErr)
	}

	if len(createInvoiceDto.TravelItems) > 0 {

		invoice.Id = types.EID(savedInvoiceDto.Id)
		travelItemList := make([]*travelItemDomain.TravelItem, 0)

		// convert travelItemDTO to TravelItem domain model

		for _, travelItemDTO := range createInvoiceDto.TravelItems {
			travelItem := travelItemDomain.NewTravelItemBuilder().
				SetTotalPrice(travelItemDTO.TotalPrice).
				SetIdInvoice(invoice.Id).
				SetId(types.EID(travelItemDTO.Id)).
				Build()
			travelItemList = append(travelItemList, travelItem)
		}

		domainService := invoiceDomain.NewInvoiceDomainService(application.travelItemRepository, application.imputationRepository, nil, application.invoiceRepository, application.TransactionManager)

		domainService.AddTravelItem(invoice, travelItemList)

	}

	transactionErr = application.TransactionManager.Commit()

	if transactionErr != nil {
		application.Logger.Error(fmt.Sprintf("[CreateInvoiceService] - Error commiting transaction: %v", transactionErr))
		return nil, transactionErr
	}

	return savedInvoiceDto, nil
}
