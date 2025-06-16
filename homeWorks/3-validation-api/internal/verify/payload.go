package verify

type SendRequest struct {
	Email string `json:"Email" validate:"required,email"`
}

type VerifyResponse struct {
	Success bool `json:"Success"`
}
type SendResponse struct {
	Success bool `json:"Success"`
}
