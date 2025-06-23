package main

import (
	"log"
	"net/http"
	"temp/configs"
	"temp/internal/auth"
	"temp/pkg/db"
)

func main() {
	cfg := configs.LoadConfig()
	_,err := db.NewDB(cfg)
	if err != nil{
		log.Fatalln("Error, not connect to db,",err)
	}
	router := http.NewServeMux()
	auth.NewAuthHandler(router, &auth.AuthHandlerDeps{
		Config: &cfg,
	})
	server := http.Server{
		Addr:    cfg.WServer.Addr,
		Handler: router,
	}
	server.ListenAndServe()

}
