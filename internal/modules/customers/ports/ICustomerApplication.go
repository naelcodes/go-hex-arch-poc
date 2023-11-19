package ports

type ICustomerApplication interface {
	CreateCustomerService()
	UpdateCustomerService()
	GetAllCountriesService()
	GetAllCustomersService()
	GetCustomerService()
	DeleteCustomerService()
}
