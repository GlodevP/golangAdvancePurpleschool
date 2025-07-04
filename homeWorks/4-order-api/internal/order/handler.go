package order

import (
	"4-order-api/config"
	"4-order-api/pkg/request"
	"4-order-api/pkg/response"
	"log"
	"net/http"
)

type OrderHandler struct {
	Dependens *OrderDependens
}

type OrderDependens struct {
	cfg *config.Config
	db  *Repository
}

func NewOrderHandle(cfg *config.Config, r *http.ServeMux, db *Repository) {
	h := OrderHandler{
		Dependens: &OrderDependens{
			cfg: cfg,
			db:  db,
		},
	}
	r.HandleFunc("GET /order/{id}", h.getOrderHandler())
	r.HandleFunc("POST /order", h.addOrderHandler())
}

func (handler OrderHandler) getOrderHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		o, err := handler.Dependens.db.GetProduct(id)
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

func (handler OrderHandler) addOrderHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := request.HandleBody[AddProductRequest](&w, r)
		if err != nil {
			log.Println(err)
			response.Json(w, AddProductResponce{Success: false}, http.StatusBadRequest)
			return
		}
		err = handler.Dependens.db.AddProduct(&Product{
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
