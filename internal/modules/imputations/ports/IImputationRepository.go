package ports

type IImputationRepository interface {
	FindByInvoiceId()
	Save()
	Update()
}
