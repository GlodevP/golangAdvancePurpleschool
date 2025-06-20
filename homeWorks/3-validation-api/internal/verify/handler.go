package verify

import (
	"3-validation-api/configs"
	"3-validation-api/internal/store"
	"3-validation-api/pkg/request"
	"3-validation-api/pkg/response"
	"net/http"
)

type VerifyHandler struct {
	dependens *VerifyDependens
}

type VerifyDependens struct {
	cfg *configs.Config
	DB  *store.DB
}

func NewVerifyHandler(cfg *configs.Config, router *http.ServeMux) error {
	db, err := store.NewDB(cfg.NameDB)
	if err != nil {
		return err
	}
	verify := VerifyHandler{
		dependens: &VerifyDependens{
			cfg: cfg,
			DB:  db,
		},
	}
	router.HandleFunc("POST /verify/send", verify.getSendHandleFunc())
	router.HandleFunc("GET /verify/{hash}", verify.getVerifyHandleFunc())
	return nil
}

func (handler *VerifyHandler) getSendHandleFunc() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		sreq, err := request.HandleBody[SendRequest](&w, r)
		if err != nil {
			res := &SendResponse{
				Success: false,
			}
			response.Json(w, res, http.StatusInternalServerError)
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
