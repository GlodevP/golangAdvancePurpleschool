package main

import (
	"4-order-api/config"
	"4-order-api/internal/order"
	"4-order-api/pkg/middleware"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func init(){
	log.SetFormatter(&log.JSONFormatter{})
}

func main() {
	cfg := config.NewConfigs()
	router := http.NewServeMux()
	db, err := order.NewRepository(*cfg)
	if err != nil {
		log.Fatalln(err)
	}
	order.NewOrderHandle(cfg, router, db)
	http.ListenAndServe(cfg.Webserver.Addr, middleware.Logging(router))

}
