package common

var (
	INVALID_MONETARY_VALUE = GenericError{nil, "monetary value can't be less than zero"}
	INVALID_CURRENCY       = GenericError{nil, "monetary operations can't be done with monetary value of different currencies"}
)

type GenericError struct {
	code    *int
	message string
}

func (g GenericError) Error() string {
	return g.message
}
