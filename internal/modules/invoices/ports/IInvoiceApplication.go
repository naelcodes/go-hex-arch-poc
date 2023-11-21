package ports

type IInvoiceApplication interface {
	GetInvoiceService()
	GetAllInvoiceService()
	GetAllTravelItemsService()
	CreateInvoiceService()
	UpdateInvoiceService()
	DeleteInvoiceService()
}
