package verify

import (
	"3-validation-api/configs"
	"3-validation-api/pkg/request"
	"3-validation-api/pkg/response"
	"net/http"
)

type VerifyHandler struct {
	dependens *VerifyDependens
}

type VerifyDependens struct {
	cfg *configs.Config
	db  []string
}

func NewVerifyHandler(cfg *configs.Config, router *http.ServeMux) {
	verify := VerifyHandler{
		dependens: &VerifyDependens{
			cfg: cfg,
			db:  []string{},
		},
	}
	router.HandleFunc("POST /verify/send", verify.getSendHandleFunc())
	router.HandleFunc("GET /verify/{hash}", verify.getVerifyHandleFunc())
}

func (handler *VerifyHandler) getSendHandleFunc() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		sreq, err := request.HandleBody[SendRequest](&w, r)
		if err != nil {
			return
		}
		err = handler.sendEmailVerify(sreq.Email)
		if err != nil {
			res := &SendResponse{
				Success: false,
			}
			response.Json(w, res, http.StatusInternalServerError)
			return
		}
		res := &SendResponse{
			Success: true,
		}
		response.Json(w, res, http.StatusOK)
	}
}

func (handler *VerifyHandler) getVerifyHandleFunc() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		hash := r.PathValue("hash")
		if handler.isHashVerify(hash) {
			res := &VerifyResponse{
				Success: true,
			}
			response.Json(w, res, http.StatusOK)
		} else {
			res := &VerifyResponse{
				Success: false,
			}
			response.Json(w, res, http.StatusBadRequest)
		}
	}
}
