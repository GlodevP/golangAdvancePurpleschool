package auth

import (
	"net/http"
	"temp/configs"
	"temp/pkg/request"
	"temp/pkg/response"
)

type AuthHandlerDeps struct {
	*configs.Config
}

type authHandler struct {
	*configs.Config
}

func NewAuthHandler(router *http.ServeMux, deps *AuthHandlerDeps) {
	authHandler := authHandler{
		Config: deps.Config,
	}

	router.HandleFunc("POST /auth/login", authHandler.getLoginHandleFunction())
	router.HandleFunc("POST /auth/register", authHandler.getRegisterHandleFunction())
}

func (handler *authHandler) getLoginHandleFunction() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := request.HandleBody[LoginRequest](&w, r)
		if err != nil {
			println(err.Error())
			return
		}
		res := LoginResponse{
			Token: "123",
		}
		response.Json(w, res, 200)
	}
}

func (handler *authHandler) getRegisterHandleFunction() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := request.HandleBody[RegisterRequest](&w, r)
		if err != nil {
			println(err.Error())
			return
		}
		res := RegisterResponse{
			Token: "123",
		}
		response.Json(w, res, 200)
	}
}
