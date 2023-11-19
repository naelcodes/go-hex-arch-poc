package postgresAdapter

type Tabler interface {
	TableName() string
}

type CustomerModel struct {
	Customer_name     string
	State             string
	Account_number    string
	Slug              int64
	Id_country        uint
	Alias             string
	Ab_key            string
	Tmc_client_number string
	Tag               string `default:"3"`
}

type CountryModel struct {
	Id   uint
	Code string
	Name string
}

func (CustomerModel) TableName() string {
	return "customer"
}

func (CountryModel) TableName() string {
	return "country"
}
