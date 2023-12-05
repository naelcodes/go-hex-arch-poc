package schema

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"strconv"

	"strings"

	"entgo.io/ent/schema/field"
)

type Money struct {
	CurrencyPrefix string
}

// Value implements the TypeValueScanner.Value method.
func (m Money) Value(amount float64) (driver.Value, error) {
	return m.CurrencyPrefix + strconv.FormatFloat(amount, 'f', 'f', 64), nil
}

// ScanValue implements the TypeValueScanner.ScanValue method.
func (m Money) ScanValue() field.ValueScanner {
	return &sql.NullString{}
}

// FromValue implements the TypeValueScanner.FromValue method.
func (m Money) FromValue(dbValue driver.Value) (float64, error) {
	moneyValue, ok := dbValue.(*sql.NullString)
	if !ok {
		return 0, fmt.Errorf("unexpected input for FromValue: %T", moneyValue)
	}
	if !moneyValue.Valid {
		return 0, nil
	}

	amountStr := moneyValue.String[1:]
	amountStr = strings.Replace(amountStr, ",", "", -1)
	amountFloat64, err := strconv.ParseFloat(amountStr, 64)

	if err != nil {
		return 0, err
	}

	return amountFloat64, nil
}
