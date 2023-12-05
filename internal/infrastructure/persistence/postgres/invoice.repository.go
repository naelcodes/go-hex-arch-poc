package postgres

import (
	"context"
	"fmt"

	"github.com/naelcodes/ab-backend/ent"
	"github.com/naelcodes/ab-backend/ent/customer"
	"github.com/naelcodes/ab-backend/ent/invoice"
	invoiceDomain "github.com/naelcodes/ab-backend/internal/core/domains/invoice-domain"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	CustomErrors "github.com/naelcodes/ab-backend/pkg/errors"
	"github.com/naelcodes/ab-backend/pkg/logger"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type InvoiceRepository struct {
	Database *ent.Client
	Context  context.Context
	Logger   *logger.Logger
}

func (repo *InvoiceRepository) Count() (*int, error) {
	return nil, nil
}
func (repo *InvoiceRepository) CountByCustomerId(customerId types.EID) (*int, error) {

	repo.Logger.Info(fmt.Sprintf("[InvoiceRepository - CountByCustomerId] Customer ID: %v", customerId))

	totalRowCount, err := repo.Database.Invoice.Query().
		Where(invoice.And(
			invoice.HasCustomerWith(customer.IDEQ(int(customerId))),
			invoice.TagEQ(invoice.Tag3))).
		Count(repo.Context)

	if err != nil {
		repo.Logger.Error(fmt.Sprintf("[InvoiceRepository - CountByCustomerId] Error counting customer's invoices: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error counting customer's invoices: %v", err))

	}

	repo.Logger.Info(fmt.Sprintf("[InvoiceRepository - CountByCustomerId] Total number of customer's invoices: %v", totalRowCount))

	return &totalRowCount, nil
}
func (repo *InvoiceRepository) GetByCustomerID(types.EID, *types.GetQueryParams, *bool) ([]*dto.GetInvoiceDTO, error) {
	return nil, nil
}

func (repo *InvoiceRepository) GetById(id types.EID) (*dto.GetInvoiceDTO, error) {
	return nil, nil
}

func (repo *InvoiceRepository) GetAll(*types.GetQueryParams) ([]*dto.GetInvoiceDTO, error) {
	return nil, nil
}

func (repo *InvoiceRepository) Save(transaction *ent.Tx, invoice *invoiceDomain.Invoice) (*dto.GetInvoiceDTO, error) {
	return nil, nil
}

func (repo *InvoiceRepository) SaveImputation(transaction *ent.Tx, invoiceEntity *invoiceDomain.Invoice) {
	repo.Logger.Info(fmt.Sprintf("[InvoiceRepository - SaveImputation] Invoice ID: %v", invoiceEntity.Id))

	updatedInvoice, err := transaction.Invoice.UpdateOneID(int(invoiceEntity.Id)).
		SetBalance(invoiceEntity.Balance).
		SetCreditApply(invoiceEntity.Credit_apply).
		SetStatus(invoice.Status(invoiceEntity.Status)).
		Save(repo.Context)

	if err != nil {
		repo.Logger.Error(fmt.Sprintf("[InvoiceRepository - SaveImputation] Error updating invoice: %v", err))
		panic(CustomErrors.RepositoryError(fmt.Errorf("error updating invoice: %v", err)))
	}

	repo.Logger.Info(fmt.Sprintf("[InvoiceRepository - SaveImputation] Updated invoice: %v", updatedInvoice))
}

func (repo *InvoiceRepository) Update(transaction *ent.Tx, invoice *invoiceDomain.Invoice) error {
	return nil
}

func (repo *InvoiceRepository) Void(id types.EID) error {
	return nil
}
