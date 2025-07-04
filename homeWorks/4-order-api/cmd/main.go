package main

import (
	"4-order-api/config"
	"4-order-api/internal/order"
	"log"
	"net/http"
)

func main() {
	cfg := config.NewConfigs()
	router := http.NewServeMux()
	db,err := order.NewRepository(*cfg)
	if err != nil {
		log.Fatalln(err)
	}
	order.NewOrderHandle(cfg, router, db)
	http.ListenAndServe(cfg.Webserver.Addr, router)

}
