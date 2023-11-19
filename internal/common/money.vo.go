package common

type Money struct {
	currency Currency
	value    float32
}

func (m *Money) Init(currency string, amount float32) error {
	if amount < 0 {
		return INVALID_MONETARY_VALUE
	}

	m.currency = Currency(currency)
	m.value = amount
	return nil
}

func (m *Money) checkCurrency(_m Money) error {
	if m.currency != _m.currency {
		return INVALID_CURRENCY
	}
	return nil
}

func (m *Money) Add(amount Money) error {
	if err := m.checkCurrency(amount); err != nil {
		return err
	}
	m.value += amount.value
	return nil
}

func (m *Money) Deduct(amount Money) error {
	if m.value == 0 {
		return INVALID_MONETARY_VALUE
	}

	m.value -= amount.value
	return nil

}
