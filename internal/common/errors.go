package common

type GenericError struct {
	code    *int
	message string
}

func (g GenericError) Error() string {
	return g.message
}
