package dto

import "github.com/naelcodes/ab-backend/internal/common"

type GetCountriesDTO struct {
	Id   int    `json:"id"`
	Code string `json:"code"`
	Name string `json:"name"`
}

type GetCustomerDTO struct {
	Id                uint   ` json:"id"`
	Customer_name     string `json:"customerName"`
	State             string `json:"state"`
	Account_number    string `json:"accountNumber"`
	Slug              uint64 `json:"slug"`
	Id_country        uint   `json:"idCountry"`
	Alias             string `json:"alias"`
	Ab_key            string `json:"abKey"`
	Tmc_client_number string `json:"tmcClientNumber"`
}

type GetCustomersDTO common.GetAllDTO[[]GetCustomerDTO]
