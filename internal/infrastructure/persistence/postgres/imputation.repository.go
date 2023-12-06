package postgres

import (
	"context"
	"fmt"

	"github.com/naelcodes/ab-backend/ent"
	imputationDomain "github.com/naelcodes/ab-backend/internal/core/domains/imputation-domain"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	CustomErrors "github.com/naelcodes/ab-backend/pkg/errors"
	"github.com/naelcodes/ab-backend/pkg/types"
	"github.com/naelcodes/ab-backend/pkg/utils"
)

type ImputationRepository struct {
	Database *ent.Client
	Context  context.Context
}

func (repo *ImputationRepository) CountByInvoiceId(idInvoice types.EID) (*int, error) {
	return nil, nil
}

func (repo *ImputationRepository) CountByPaymentId(idPayment types.EID) (*int, error) {
	return nil, nil
}

func (repo *ImputationRepository) GetByInvoiceId(idInvoice types.EID, queryParams *types.GetQueryParams) ([]*dto.GetImputationDetails, error) {
	return nil, nil
}

func (repo *ImputationRepository) GetByPaymentAndInvoiceId(idPayment types.EID, idInvoice types.EID) (*dto.ImputationDetails, error) {
	return nil, nil
}

func (repo *ImputationRepository) Update(transaction *ent.Tx, imputationEntity *imputationDomain.Imputation) error {
	return nil
}

func (repo *ImputationRepository) SaveAll(transaction *ent.Tx, imputationEntities []*imputationDomain.Imputation) {

	utils.Logger.Info(fmt.Sprintf("[ImputationRepository - SaveAll] Saving %v imputations", len(imputationEntities)))

	for _, i := range imputationEntities {
		utils.Logger.Info(fmt.Sprintf("[ImputationRepository - SaveAll] Saving imputation: %v", i))

		updatedImputation, err := transaction.Imputation.UpdateOneID(int(i.Id)).
			SetInvoiceID(int(i.IdInvoice)).
			SetPaymentID(int(i.IdPayment)).
			SetAmountApply(i.AmountApplied).
			Save(repo.Context)

		if err != nil {
			utils.Logger.Error(fmt.Sprintf("[ImputationRepository - SaveAll] Error updating imputation: %v", err))
			panic(CustomErrors.RepositoryError(fmt.Errorf("error updating imputation: %v", err)))
		}
		utils.Logger.Info(fmt.Sprintf("[ImputationRepository - SaveAll] Updated imputation: %v", updatedImputation))

	}

	utils.Logger.Info("[ImputationRepository - SaveAll] Saved imputations")
}

func (repo *ImputationRepository) Delete(transaction *ent.Tx, id types.EID) error {
	return nil
}
