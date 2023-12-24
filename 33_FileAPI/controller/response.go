package controller

type ErrorResponse struct {
	ErrorDescription string `json:"error_descriptin"`
}
type UploadFileResponse struct {
	FileName string `json:"file_name"`
}
