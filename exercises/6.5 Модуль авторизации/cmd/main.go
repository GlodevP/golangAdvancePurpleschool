package main

import (
	"net/http"
	"temp/configs"
	"temp/internal/auth"
)

func main() {
	cfg := configs.LoadConfig()
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
