package order

import (
	"4-order-api/config"
	"4-order-api/pkg/helpers"
	"4-order-api/pkg/request"
	"4-order-api/pkg/response"
	"log"
	"net/http"

	"gorm.io/gorm"
)

type OrderHandler struct {
	Dependens *OrderDependens
}

type OrderDependens struct {
	cfg *config.Config
	repository  *Repository
}

func NewOrderHandle(cfg *config.Config, r *http.ServeMux, db *Repository) {
	h := OrderHandler{
		Dependens: &OrderDependens{
			cfg: cfg,
			repository:  db,
		},
	}
	r.HandleFunc("POST /order", h.addOrderHandler())
	r.HandleFunc("GET /order/{id}", h.getOrderHandler())
	r.HandleFunc("PATCH /order/{id}", h.updateOrderHandler())
	r.HandleFunc("DELETE /order/{id}", h.deletOrderHandler())
}

func (handler OrderHandler) addOrderHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := request.HandleBody[AddProductRequest](&w, r)
		if err != nil {
			log.Println(err)
			response.Json(w, AddProductResponce{Success: false}, http.StatusBadRequest)
			return
		}
		err = handler.Dependens.repository.Add(&Product{
			Name:        req.Name,
			Description: req.Description,
		})
		if err != nil {
			log.Println(err)
			response.Json(w, AddProductResponce{Success: false}, http.StatusBadRequest)
			return
		}

		response.Json(w, AddProductResponce{Success: true}, http.StatusOK)
	}
}

func (handler OrderHandler) getOrderHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		u,err := helpers.StingToUint(id)
		if err != nil {
			log.Println(err)
			response.Json(w, GetProductResponce{
				Success: false,
			}, http.StatusBadRequest)
			return
		}
		o, err := handler.Dependens.repository.GetByID(u)
		if err != nil {
			log.Println(err)
			response.Json(w, GetProductResponce{
				Success: false,
			}, http.StatusBadRequest)
			return
		}
		response.Json(w, &GetProductResponce{
			Success: true,
			Product:   *o,
		}, http.StatusOK)
	}
}


func (handler OrderHandler) updateOrderHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		u,err := helpers.StingToUint(id)
		if err != nil {
			log.Println(err)
			response.Json(w, UpdateResponse{
				Success: false,
			}, http.StatusBadRequest)
			return
		}
		body,err := request.HandleBody[UpdateRequest](&w,r)
		if err != nil {
			log.Println(err)
			response.Json(w, UpdateResponse{
				Success: false,
			}, http.StatusBadRequest)
			return
		}

		o, err := handler.Dependens.repository.Update(&Product{
			Model: gorm.Model{
				ID: u,
			},
			Name: body.Name,
			Description: body.Description,
		})
		if err != nil {
			log.Println(err)
			response.Json(w, UpdateResponse{
				Success: false,
			}, http.StatusBadRequest)
			return
		}
		response.Json(w, UpdateResponse{
			Success: true,
			Product:   *o,
		}, http.StatusOK)
	}
}

func (handler OrderHandler) deletOrderHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		u,err := helpers.StingToUint(id)
		if err != nil {
			log.Println(err)
			response.Json(w, DeletResponse{
				Success: false,
			}, http.StatusBadRequest)
			return
		}
		err = handler.Dependens.repository.Delete(u)
		if err != nil {
			log.Println(err)
			response.Json(w, DeletResponse{
				Success: false,
			}, http.StatusBadRequest)
			return
		}
		response.Json(w, DeletResponse{
			Success: true,
		}, http.StatusOK)
	}
}
