package schemas

type ErrorResponse struct {
	Message string `json:"message"`
}

type ValidationErrorResponse struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}
