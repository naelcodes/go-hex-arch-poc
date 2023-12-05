package imputationDomain

import (
	"github.com/naelcodes/ab-backend/ent"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type IImputationRepository interface {
	CountByInvoiceId(idInvoice types.EID) (*int, error)
	CountByPaymentId(idPayment types.EID) (*int, error)
	GetByInvoiceId(idInvoice types.EID, queryParams *types.GetQueryParams) ([]*dto.GetImputationDetails, error)
	GetByPaymentAndInvoiceId(idPayment types.EID, idInvoice types.EID) (*dto.ImputationDetails, error)

	Update(*ent.Tx, *Imputation) error
	SaveAll(*ent.Tx, []*Imputation)

	Delete(*ent.Tx, types.EID) error
}
