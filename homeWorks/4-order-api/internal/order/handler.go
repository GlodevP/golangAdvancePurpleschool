package order

import (
	"4-order-api/config"
	"4-order-api/internal/store"
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
	db  *store.DB
}

func NewOrderHandle(cfg *config.Config, r *http.ServeMux, db *store.DB) {
	h := OrderHandler{
		Dependens: &OrderDependens{
			cfg: cfg,
			db:  db,
		},
	}
	r.HandleFunc("GET /order/{id}", h.getOrderHandler())
	r.HandleFunc("POST /order/add", h.addOrderHandler())
}

func (handler OrderHandler) getOrderHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		o, err := handler.Dependens.db.GetOrder(id)
		if err != nil {
			log.Println(err)
			response.Json(w, GetOrderResponce{
				Success: false,
			}, http.StatusBadRequest)
			return
		}
		response.Json(w, &GetOrderResponce{
			Success: true,
			Order:   *o,
		}, http.StatusOK)
	}
}

func (handler OrderHandler) addOrderHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		req, err := request.HandleBody[AddOrderRequest](&w, r)
		if err != nil {
			log.Println(err)
			response.Json(w, AddOrderResponce{Success: false}, http.StatusBadRequest)
			return
		}
		err = handler.Dependens.db.AddOrder(&store.Order{
			Name:        req.Name,
			Description: req.Description,
		})
		if err != nil {
			log.Println(err)
			response.Json(w, AddOrderResponce{Success: false}, http.StatusBadRequest)
			return
		}

		response.Json(w, AddOrderResponce{Success: true}, http.StatusOK)
		return

	}
}
