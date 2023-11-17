package models

type Customer struct {
	Id                uint   `gorm:"primaryKey" json:"id"`
	Customer_name     string `json:"customerName"`
	State             string `json:"state"`
	Account_number    string `json:"accountNumber"`
	Slug              int64  `json:"slug"`            //needed due to constraint - unique (add a random value before each save)
	Id_currency       uint   `json:"idCurrency"`      //needed due to constraint
	Id_country        uint   `json:"idCountry"`       //needed due to constraint
	Alias             string `json:"alias"`           //needed due to constraint - unique
	Ab_key            string `json:"abKey"`           //needed due to constraint - unique
	Tmc_client_number string `json:"tmcClientNumber"` // needed due to constraint unique

}

type Tabler interface {
	TableName() string
}

func (Customer) TableName() string {
	return "customer"
}
