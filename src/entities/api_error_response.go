package entities

type ApiErrorResponse struct {
	Error       string `json:"error"`
	Description string `json:"description"`
}
