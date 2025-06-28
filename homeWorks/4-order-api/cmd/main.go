package main

import (
	"4-order-api/config"
	"4-order-api/internal/order"
	"4-order-api/internal/store"
	"net/http"
)

func main(){
	cfg := config.NewConfigs()
	router := http.NewServeMux()
	db := store.NewDB(*cfg)
	order.NewOrderHandle(cfg,router,db)
	http.ListenAndServe(cfg.Webserver.Addr,router)
	
}