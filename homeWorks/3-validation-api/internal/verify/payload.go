package verify

type VerifyResponse struct{
	Success bool `json: "Success"`
}
type SendResponse struct{
	Success bool `json: "Success"`
	Email string `json: "Email"`
}