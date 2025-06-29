package main

import (
	"4-order-api/config"
	"4-order-api/internal/order"
	"4-order-api/internal/store"
	"log"
	"net/http"
)

func main() {
	cfg := config.NewConfigs()
	router := http.NewServeMux()
	db,err := store.NewDB(*cfg)
	if err != nil {
		log.Fatalln(err)
	}
	order.NewOrderHandle(cfg, router, db)
	http.ListenAndServe(cfg.Webserver.Addr, router)

}
