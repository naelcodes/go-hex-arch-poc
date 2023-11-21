package ports

type IPaymentApplication interface {
	AddCustomerPaymentService()
	UpdatePaymentService()
	GetAllPaymentService()
	GetCustomerPaymentService()
	DeletePaymentService()
}
