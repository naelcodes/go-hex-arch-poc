package paymentDomain

import "github.com/naelcodes/ab-backend/pkg/types"

type IPaymentRepository interface {
	CountByCustomerID(customerId types.EID) (int, error)
}
