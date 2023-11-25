package common

import "errors"

type Currency string
type Money struct {
	currency Currency
	value    float32
}

var (
	INVALID_MONETARY_VALUE = errors.New("monetary value can't be less than zero")
	INVALID_CURRENCY       = errors.New("monetary operations can't be done with monetary value of different currencies")
)

func (m *Money) New(currency string, amount float32) error {
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
