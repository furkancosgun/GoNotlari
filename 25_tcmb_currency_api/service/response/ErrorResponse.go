package response

type ErrorResponse struct {
	ErrorDescription string
}

func NewErrorResponseWithError(err error) ErrorResponse {
	return NewErrorResponseWithString(err.Error())
}
func NewErrorResponseWithString(err string) ErrorResponse {
	return ErrorResponse{ErrorDescription: err}
}
