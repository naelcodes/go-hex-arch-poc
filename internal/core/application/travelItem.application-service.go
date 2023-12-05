package application

import (
	"github.com/naelcodes/ab-backend/internal/core/dto"
	"github.com/naelcodes/ab-backend/pkg/types"
)

func (application *Application) GetAllTravelItemService(queryParams *types.GetQueryParams) (*dto.GetAllTravelItemDTO, error) {

	totalRowCount, err := application.travelItemRepository.Count()

	if err != nil {
		return nil, err
	}

	if queryParams == nil || (queryParams.PageNumber == nil && queryParams.PageSize == nil) {
		if queryParams == nil {
			queryParams = new(types.GetQueryParams)
		}
		queryParams.PageNumber = new(int)
		queryParams.PageSize = new(int)
		*queryParams.PageNumber = 0
		*queryParams.PageSize = *totalRowCount
	}

	travelItems, err := application.travelItemRepository.GetAll(queryParams)

	if err != nil {
		return nil, err
	}

	getTravelItemDTO := &dto.GetAllTravelItemDTO{
		Data:          travelItems,
		PageNumber:    *queryParams.PageNumber,
		PageSize:      *queryParams.PageSize,
		TotalRowCount: *totalRowCount,
	}

	return getTravelItemDTO, nil

}
