package auth

type LoginRequest struct {
	Email    string `json:"Email" validate:"required,email"`
	Password string `json:"Password" validate:"required"`
}

type RegisterRequest struct {
	Name     string `json:"Name" validate:"required"`
	Email    string `json:"Email" validate:"required,email"`
	Password string `json:"Password" validate:"required"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterResponse struct {
	Token string `json:"token"`
}
