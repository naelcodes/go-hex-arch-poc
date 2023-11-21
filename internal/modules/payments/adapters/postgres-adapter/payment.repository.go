package postgresAdapter

type PaymentRepository struct {
	Database any
}

func (repo *PaymentRepository) GetById(id uint) (*PaymentModel, error) {
	return nil, nil
}

func (repo *PaymentRepository) Save(payment *PaymentModel) error {
	return nil
}
