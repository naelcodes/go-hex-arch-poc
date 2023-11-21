package application

import "github.com/naelcodes/ab-backend/internal/modules/imputations/ports"

type ImputationApplication struct {
	repository ports.IImputationRepository
}

func (application *ImputationApplication) Init(imputationRepository ports.IImputationRepository) {
	application.repository = imputationRepository
}
