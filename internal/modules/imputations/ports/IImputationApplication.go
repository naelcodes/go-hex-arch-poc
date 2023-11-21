package ports

type IImputationApplication interface {
	ApplyImputationsService()
	GetInvoiceImputationsService()
	UpdateInvoiceImputationsService()
}
