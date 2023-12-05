package postgres

import (
	"context"
	"fmt"

	"github.com/naelcodes/ab-backend/ent"
	"github.com/naelcodes/ab-backend/ent/invoice"
	"github.com/naelcodes/ab-backend/ent/travelitem"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	CustomErrors "github.com/naelcodes/ab-backend/pkg/errors"
	"github.com/naelcodes/ab-backend/pkg/logger"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type TravelItemRepository struct {
	Database *ent.Client
	Context  context.Context
	Logger   *logger.Logger
}

func (repo *TravelItemRepository) Count() (*int, error) {
	repo.Logger.Info("[TravelItemRepository - Count] counting travel items")
	count, err := repo.Database.TravelItem.Query().
		Where(
			travelitem.And(
				travelitem.TransactionTypeEQ("sales"),
				travelitem.StatusEQ("pending"),
				travelitem.ProductTypeEQ("flight"),
				travelitem.Not(travelitem.HasInvoice()))).Count(repo.Context)
	if err != nil {
		repo.Logger.Error(fmt.Sprintf("[TravelItemRepository - Count] Error counting travel items: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error counting travel items: %v", err))
	}

	repo.Logger.Info(fmt.Sprintf("[TravelItemRepository - Count] Total number of travel items: %v", count))
	return &count, nil
}

func (repo *TravelItemRepository) GetAll(queryParams *types.GetQueryParams) ([]*dto.TravelItemDTO, error) {

	repo.Logger.Info("[TravelItemRepository - GetAll] Getting travel items")

	TravelItemQuery := repo.Database.TravelItem.Query().
		Where(
			travelitem.And(
				travelitem.TransactionTypeEQ("sales"),
				travelitem.StatusEQ("pending"),
				travelitem.ProductTypeEQ("flight"),
				travelitem.Not(travelitem.HasInvoice()))).
		Select(
			travelitem.FieldID,
			travelitem.FieldTotalPrice,
			travelitem.FieldItinerary,
			travelitem.FieldTravelerName,
			travelitem.FieldTicketNumber)

	if queryParams != nil && queryParams.PageNumber != nil && queryParams.PageSize != nil {
		TravelItemQuery.Offset(*queryParams.PageNumber * *queryParams.PageSize).Limit(*queryParams.PageSize)
	}

	travelItems, err := TravelItemQuery.All(repo.Context)

	if err != nil {
		repo.Logger.Error(fmt.Sprintf("[TravelItemRepository - GetAll] Error getting travel items: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error getting travel items: %v", err))
	}

	repo.Logger.Info(fmt.Sprintf("[TravelItemRepository - GetAll] Found %v travel items", len(travelItems)))
	travelItemDTOList := TravelItemModelListToDTOList(travelItems)

	return travelItemDTOList, nil

}

func (repo *TravelItemRepository) UpdateByInvoiceId(transaction *ent.Tx, invoiceId *types.EID, travelItemIds []int) error {

	repo.Logger.Info(fmt.Sprintf("[TravelItemRepository - UpdateByInvoiceId] Invoice ID: %v in  %v travel items", invoiceId, len(travelItemIds)))

	UpdateQuery := transaction.TravelItem.Update().Where(travelitem.IDIn(travelItemIds...))

	if invoiceId != nil {
		UpdateQuery.SetInvoiceID(int(*invoiceId)).SetStatus(travelitem.StatusInvoiced)
	} else {
		UpdateQuery.ClearInvoice().SetStatus(travelitem.StatusPending)
	}

	affectedRows, err := UpdateQuery.Save(repo.Context)

	if err != nil {
		repo.Logger.Error(fmt.Sprintf("[TravelItemRepository - UpdateByInvoiceId] Error updating travel item: %v", err))
		return CustomErrors.RepositoryError(fmt.Errorf("error updating travel item: %v", err))
	}

	repo.Logger.Info(fmt.Sprintf("[TravelItemRepository - UpdateByInvoiceId] Updated %v travel items", affectedRows))
	return nil
}

func (repo *TravelItemRepository) GetByInvoiceId(invoiceId types.EID) ([]*dto.TravelItemDTO, error) {

	repo.Logger.Info(fmt.Sprintf("[TravelItemRepository - GetByInvoiceId] Invoice ID: %v", invoiceId))

	TravelItems, err := repo.Database.TravelItem.Query().
		Where(travelitem.HasInvoiceWith(
			invoice.IDEQ(int(invoiceId)))).
		Select(travelitem.FieldID).All(repo.Context)

	if err != nil {
		repo.Logger.Error(fmt.Sprintf("[TravelItemRepository - GetByInvoiceId] Error getting travel items: %v", err))
		return nil, CustomErrors.RepositoryError(fmt.Errorf("error getting travel items: %v", err))
	}

	repo.Logger.Info(fmt.Sprintf("[TravelItemRepository - GetByInvoiceId] Found %v travel items", len(TravelItems)))

	travelItemDTOList := TravelItemModelListToDTOList(TravelItems)

	return travelItemDTOList, nil
}
