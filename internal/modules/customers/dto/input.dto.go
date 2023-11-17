package dto

type CreateCustomerDTO struct {
	Customer_name     string `json:"customerName"`
	State             string `json:"state"`
	Account_number    string `json:"accountNumber"`
	Id_country        uint   `json:"idCountry"`
	Alias             string `json:"alias"`
	Tmc_client_number string `json:"tmcClientNumber"`
}

type UpdateCustomerDTO struct {
	Customer_name     *string `json:"customerName"`
	State             *string `json:"state"`
	Account_number    *string `json:"accountNumber"`
	Id_country        *uint   `json:"idCountry"`
	Alias             *string `json:"alias"`
	Tmc_client_number *string `json:"tmcClientNumber"`
}
