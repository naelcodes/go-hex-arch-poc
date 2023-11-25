package common

// var (
//
//	InvalidRequest = GenericErrorhandlerResponse{code: 400, success: false, message: "Invalid Request"}
//	NotFound       = GenericErrorhandlerResponse{code: 404, success: false, message: "Not Found"}
//	Forbidden      = GenericErrorhandlerResponse{code: 403, success: false, message: "Forbidden"}
//	Conflict       = GenericErrorhandlerResponse{code: 409, success: false, message: "Conflict"}
//	ServerError    = GenericErrorhandlerResponse{code: 500, success: false, message: "Internal Server Error"}
//
// )
type GenericErrorhandlerResponse struct {
	code    int
	success bool
	message string
}

func (g GenericErrorhandlerResponse) Error() string {
	return g.message
}
