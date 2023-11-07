package models

type Customer struct {
	Id              uint    `gorm:"primaryKey" json:"id"`
	Customer_name   string  `json:"customerName"`
	Street          string  `json:"street"`
	City            string  `json:"city"`
	State           string  `json:"state"`
	Zip_code        string  `json:"zipCode"`
	Notes           string  `json:"notes"`
	Terms           int     `json:"terms"`
	Account_number  string  `json:"accountNumber"`
	// Tax_id          string  `json:"taxId"`
	// Balance         string  `json:"balance" sql:"type:decimal(10,10)"`
	// Credit_limit    float64 `json:"creditLimit"`
	Is_active       bool    `json:"isActive"`
	Is_sub_agency   bool    `json:"isSubAgency"`
	// Opening_balance float64 `json:"openingBalance"`
	Language        string  `json:"language"`
	Slug            int32   `json:"slug"`       //needed due to constraint - unique (add a random value before each save)
	Id_currency     uint    `json:"idCurrency"` //needed due to constraint
	Id_country      uint    `json:"idCountry"`  //needed due to constraint
	// Irs_share_key   string  `json:"irsShareKey"`
	// Currency_rate   float64 `json:"currencyRate"`
	Agency          string  `json:"agency"`
	// Opening_balance_date  time.Time `json:"openingBalanceDate"`
	// Avoid_deletion    bool   `json:"avoidDeletion"`
	// Is_editable       bool   `json:"isEditable"`
	Alias             string `json:"alias"` //needed due to constraint - unique
	// Already_used      int    `json:"alreadyUsed"`
	Ab_key            string `json:"abKey"`           //needed due to constraint - unique
	Tmc_client_number string `json:"tmcClientNumber"` // needed due to constraint unique

}

type Tabler interface {
	TableName() string
}

func (Customer) TableName() string {
	return "customer"
}