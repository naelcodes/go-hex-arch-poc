package postgres

import (
	"context"
	"fmt"

	"github.com/naelcodes/ab-backend/ent"
	"github.com/naelcodes/ab-backend/ent/customer"
	"github.com/naelcodes/ab-backend/ent/invoice"
	"github.com/naelcodes/ab-backend/ent/travelitem"
	invoiceDomain "github.com/naelcodes/ab-backend/internal/core/domains/invoice-domain"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	CustomErrors "github.com/naelcodes/ab-backend/pkg/errors"
	"github.com/naelcodes/ab-backend/pkg/types"
	"github.com/naelcodes/ab-backend/pkg/utils"
)

type InvoiceRepository struct {
	Database *ent.Client
	Context  context.Context
}

func (repo *InvoiceRepository) Exists(idInvoice types.EID) (bool, error) {

	utils.Logger.Info(fmt.Sprintf("[InvoiceRepository - Exists] Invoice ID: %v", idInvoice))

	invoiceExists, err := repo.Database.Invoice.Query().Where(invoice.IDEQ(int(idInvoice))).Exist(repo.Context)
	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[InvoiceRepository - Exists] Error checking if invoice exists: %v", err))
		return false, CustomErrors.RepositoryError(fmt.Errorf("error checking if invoice exists: %v", err))
	}

	utils.Logger.Info(fmt.Sprintf("[InvoiceRepository - Exists] Invoice exists: %v", invoiceExists))
	return invoiceExists, nil
}

func (repo *InvoiceRepository) Count() (*int, error) {

	utils.Logger.Info("[InvoiceRepository - Count]")

	totalRowCount, err := repo.Database.Invoice.Query().Where(invoice.TagEQ(invoice.Tag3)).Count(repo.Context)
	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[InvoiceRepository - Count] Error counting invoices: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error counting invoices: %v", err))
	}

	utils.Logger.Info(fmt.Sprintf("[InvoiceRepository - Count] Total number of invoices: %v", totalRowCount))
	return &totalRowCount, nil
}
func (repo *InvoiceRepository) CountByCustomerId(customerId types.EID) (*int, error) {

	utils.Logger.Info(fmt.Sprintf("[InvoiceRepository - CountByCustomerId] Customer ID: %v", customerId))

	totalRowCount, err := repo.Database.Invoice.Query().
		Where(invoice.And(
			invoice.HasCustomerWith(customer.IDEQ(int(customerId))),
			invoice.TagEQ(invoice.Tag3))).
		Count(repo.Context)

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[InvoiceRepository - CountByCustomerId] Error counting customer's invoices: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error counting customer's invoices: %v", err))

	}

	utils.Logger.Info(fmt.Sprintf("[InvoiceRepository - CountByCustomerId] Total number of customer's invoices: %v", totalRowCount))

	return &totalRowCount, nil
}
func (repo *InvoiceRepository) GetByCustomerID(id types.EID, queryParams *types.GetQueryParams, paid bool) (*dto.GetCustomerInvoicesDTO, error) {

	utils.Logger.Info(fmt.Sprintf("[InvoiceRepository - GetByCustomerID] Customer ID: %v", id))
	utils.Logger.Info(fmt.Sprintf("[InvoiceRepository - GetByCustomerID] Query Params: %v", queryParams))

	invoiceQuery := repo.Database.Invoice.Query().Where(
		invoice.And(
			invoice.TagEQ(invoice.Tag3),
			invoice.HasCustomerWith(customer.IDEQ(int(id))),
		),
	)

	if paid {
		invoiceQuery.Where(invoice.StatusEQ(invoice.StatusPaid))
	} else {
		invoiceQuery.Where(invoice.StatusEQ(invoice.StatusUnpaid))

	}

	totalRowCount, err := invoiceQuery.Count(repo.Context)
	if err != nil {
		return nil, err
	}

	pageNumber := 0
	pageSize := totalRowCount

	if queryParams != nil {
		if queryParams.PageNumber != nil && queryParams.PageSize != nil {
			pageNumber = *queryParams.PageNumber
			pageSize = *queryParams.PageSize
			invoiceQuery.Offset(pageNumber * pageSize).Limit(pageSize)

		}
	}

	invoices, err := invoiceQuery.All(repo.Context)
	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[InvoiceRepository - GetByCustomerID] Error getting customer(id:%v) invoices: %v", id, err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error getting  customer(id:%v) invoices: %v", id, err))
	}

	utils.Logger.Info(fmt.Sprintf("[InvoiceRepository - GetByCustomerID] Total number of customer(id:%v) invoices: %v", id, len(invoices)))

	invoiceDTOList := InvoiceModelListToDTOList(invoices, false)
	customerInvoiceDTO := new(dto.CustomerInvoice)
	customerInvoiceDTO.IdCustomer = int(id)
	customerInvoiceDTO.Invoices = invoiceDTOList

	getCustomerInvoicesDTO := new(dto.GetCustomerInvoicesDTO)
	getCustomerInvoicesDTO.Data = customerInvoiceDTO
	getCustomerInvoicesDTO.PageNumber = pageNumber
	getCustomerInvoicesDTO.PageSize = pageSize
	getCustomerInvoicesDTO.TotalRowCount = totalRowCount

	return getCustomerInvoicesDTO, nil
}

func (repo *InvoiceRepository) GetById(id types.EID, queryParams *types.GetQueryParams) (*dto.GetInvoiceDTO, error) {

	utils.Logger.Info(fmt.Sprintf("[InvoiceRepository - GetById] Invoice ID: %v", id))
	utils.Logger.Info(fmt.Sprintf("[InvoiceRepository - GetById] Query Params: %v", queryParams))

	embedCustomer := false
	invoiceQuery := repo.Database.Invoice.Query().Where(invoice.IDEQ(int(id)))

	if queryParams != nil && queryParams.Embed != nil && *queryParams.Embed == "customer" {
		embedCustomer = true
		invoiceQuery.WithCustomer()

	} else {
		invoiceQuery.WithCustomer(func(cq *ent.CustomerQuery) {
			cq.Select(customer.FieldID)
		})
	}

	invoice, err := invoiceQuery.Only(repo.Context)

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[InvoiceRepository - GetById] Error getting invoice: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error getting invoice: %v", err))
	}

	utils.Logger.Info("[InvoiceRepository - GetById] Invoice found")

	invoiceDTO := InvoiceModelToDTO(invoice, embedCustomer)

	utils.Logger.Info(fmt.Sprintf("[InvoiceRepository - GetById] Invoice DTO: %v", invoiceDTO))

	return invoiceDTO, nil
}

func (repo *InvoiceRepository) GetAll(queryParams *types.GetQueryParams) (*dto.GetAllInvoiceDTO, error) {

	embedCustomer := false
	totalRowCount, err := repo.Count()
	if err != nil {
		return nil, err
	}

	pageNumber := 0
	pageSize := *totalRowCount

	invoiceQuery := repo.Database.Invoice.Query().WithTravelItems().Where(invoice.TagEQ(invoice.Tag3))

	utils.Logger.Info(fmt.Sprintf("[InvoiceRepository - GetAll] QueryParams: %v", queryParams))

	if queryParams != nil {

		utils.Logger.Info(fmt.Sprintf("[InvoiceRepository - GetAll] Search: %v", queryParams.Fields))

		if queryParams.Embed != nil && *queryParams.Embed == "customer" {
			embedCustomer = true
			invoiceQuery.WithCustomer()
		}

		if queryParams.PageNumber != nil && queryParams.PageSize != nil {

			pageNumber = *queryParams.PageNumber
			pageSize = *queryParams.PageSize
			invoiceQuery.Offset(pageNumber * pageSize).Limit(pageSize)
		}
	}

	invoices, err := invoiceQuery.Order(ent.Asc(invoice.FieldInvoiceNumber)).All(repo.Context)

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[InvoiceRepository - GetAll] Error getting invoices records: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error getting invoices records: %v", err))
	}

	utils.Logger.Info(fmt.Sprintf("[InvoiceRepository - GetAll] Invoices records: %v", invoices))
	invoiceDTOList := InvoiceModelListToDTOList(invoices, embedCustomer)
	utils.Logger.Info(fmt.Sprintf("[InvoiceRepository - GetAll] Invoices DTO: %v", invoiceDTOList))

	getAllInvoiceDTO := new(dto.GetAllInvoiceDTO)
	getAllInvoiceDTO.Data = invoiceDTOList
	getAllInvoiceDTO.PageNumber = pageNumber
	getAllInvoiceDTO.PageSize = pageSize
	getAllInvoiceDTO.TotalRowCount = *totalRowCount

	utils.Logger.Info(fmt.Sprintf("[InvoiceRepository - GetAll] GetAllInvoiceDTO: %v", getAllInvoiceDTO))
	return getAllInvoiceDTO, nil
}

func (repo *InvoiceRepository) Save(transaction *ent.Tx, invoiceDomainModel *invoiceDomain.Invoice) (*dto.GetInvoiceDTO, error) {

	utils.Logger.Info(fmt.Sprintf("[InvoiceRepository - Save] Invoice: %v", invoiceDomainModel))

	totalRowCount, err := repo.Count()

	if err != nil {
		return nil, err
	}

	createdInvoice, err := transaction.Invoice.Create().
		SetCreationDate(invoiceDomainModel.CreationDate).
		SetInvoiceNumber(utils.GenerateCode("inv", *totalRowCount+1)).
		SetDueDate(invoiceDomainModel.DueDate).
		SetCustomerID(int(invoiceDomainModel.IdCustomer)).
		SetAmount(invoiceDomainModel.Amount).
		SetNetAmount(invoiceDomainModel.Amount).
		SetBaseAmount(invoiceDomainModel.Amount).
		SetCreditApply(invoiceDomainModel.Credit_apply).
		SetBalance(invoiceDomainModel.Balance).
		Save(repo.Context)

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[InvoiceRepository - Save] Error creating invoice: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error creating invoice: %v", err))
	}

	utils.Logger.Info(fmt.Sprintf("[InvoiceRepository - Save] Created invoice: %v", createdInvoice))

	updatedRowCount, err := transaction.TravelItem.Update().
		Where(travelitem.IDIn(invoiceDomainModel.TravelItemsId...)).
		SetInvoiceID(int(createdInvoice.ID)).
		SetStatus(travelitem.StatusInvoiced).
		Save(repo.Context)

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[InvoiceRepository - Save] Error updating travel items: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error updating travel items: %v", err))
	}

	utils.Logger.Info(fmt.Sprintf("[InvoiceRepository - Save] Updated travel items: %v", updatedRowCount))

	return InvoiceModelToDTO(createdInvoice, false), nil
}

func (repo *InvoiceRepository) SaveImputation(transaction *ent.Tx, invoiceDomainModel *invoiceDomain.Invoice) error {
	utils.Logger.Info(fmt.Sprintf("[InvoiceRepository - SaveImputation] Invoice ID: %v", invoiceDomainModel.Id))

	updatedInvoice, err := transaction.Invoice.UpdateOneID(int(invoiceDomainModel.Id)).
		SetBalance(invoiceDomainModel.Balance).
		SetCreditApply(invoiceDomainModel.Credit_apply).
		SetStatus(invoice.Status(invoiceDomainModel.Status)).
		Save(repo.Context)

	if err != nil {
		utils.Logger.Error(fmt.Sprintf("[InvoiceRepository - SaveImputation] Error updating invoice: %v", err))
		return CustomErrors.RepositoryError(fmt.Errorf("error updating invoice: %v", err))
	}

	utils.Logger.Info(fmt.Sprintf("[InvoiceRepository - SaveImputation] Updated invoice: %v", updatedInvoice))

	return nil
}

func (repo *InvoiceRepository) Update(transaction *ent.Tx, invoice *invoiceDomain.Invoice) error {
	return nil
}

func (repo *InvoiceRepository) Void(id types.EID) error {
	return nil
}
