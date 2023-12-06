package invoiceDomain

import (
	"fmt"
	"time"

	"github.com/naelcodes/ab-backend/pkg/types"

	CustomErrors "github.com/naelcodes/ab-backend/pkg/errors"
)

type Invoice struct {
	types.BaseEntity
	CreationDate  string
	InvoiceNumber string
	Status        string
	DueDate       string
	Amount        float64
	Balance       float64
	Credit_apply  float64
	IdCustomer    types.EID
	TravelItemsId []int
}

func (i *Invoice) CheckDates() error {
	creationDate, _ := time.Parse("2006-01-02", i.CreationDate)
	dueDate, _ := time.Parse("2006-01-02", i.DueDate)
	if dueDate.Before(creationDate) {
		return CustomErrors.DomainError(fmt.Errorf("due date can't be less than creation date"))
	}
	return nil
}
func (i *Invoice) ApplyImputation(imputedAmount float64) error {
	if i.Status == "paid" {
		return CustomErrors.DomainError(fmt.Errorf("imputation can't be applied, invoice is already paid"))
	}

	i.Credit_apply += imputedAmount
	err := i.CalculateBalance()

	if err != nil {
		return err
	}
	return nil
}

func (i *Invoice) CalculateBalance() error {

	if i.Credit_apply > i.Amount {
		return CustomErrors.DomainError(fmt.Errorf("the balance of an invoice can't be less than 0. credit_apply can't be greater than  invoice amount"))
	}
	i.Balance = i.Amount - i.Credit_apply
	i.UpdateStatus()

	return nil
}

func (i *Invoice) UpdateStatus() {

	if i.Credit_apply == i.Amount && i.Balance == 0 {
		i.Status = "paid"
	} else {
		i.Status = "unpaid"
	}
}
