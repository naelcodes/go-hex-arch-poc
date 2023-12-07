package imputationDomain

import (
	"github.com/naelcodes/ab-backend/ent"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type IImputationRepository interface {
	CountByInvoiceId(idInvoice types.EID) (*int, error)
	CountByPaymentId(idPayment types.EID) (*int, error)
	GetByInvoiceId(idInvoice types.EID) (*dto.GetInvoiceImputationDTO, error)
	GetByPaymentAndInvoiceId(idPayment types.EID, idInvoice types.EID) (*bool, *ent.Imputation, error)

	Update(*ent.Tx, *Imputation) (int, error)
	SaveAll(*ent.Tx, []*Imputation) (int, error)

	Delete(*ent.Tx, types.EID) (int, error)
}
