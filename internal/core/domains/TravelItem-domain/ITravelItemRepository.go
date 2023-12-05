package travelItemDomain

import (
	"github.com/naelcodes/ab-backend/ent"
	"github.com/naelcodes/ab-backend/internal/core/dto"
	"github.com/naelcodes/ab-backend/pkg/types"
)

type ITravelItemRepository interface {
	Count() (*int, error)
	GetByInvoiceId(types.EID) ([]*dto.TravelItemDTO, error)
	GetAll(*types.GetQueryParams) ([]*dto.TravelItemDTO, error)
	UpdateByInvoiceId(*ent.Tx, *types.EID, []int) error
}
