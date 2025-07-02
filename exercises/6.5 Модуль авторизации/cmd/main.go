package main

import (
	"log"
	"net/http"
	"temp/configs"
	"temp/internal/auth"
	"temp/internal/link"
	"temp/pkg/db"
)

func main() {
	cfg := configs.LoadConfig()
	db,err := db.NewDB(cfg)
	if err != nil{
		log.Fatalln("Error, not connect to db,",err)
	}
	LinkRepository := link.NewLinkRepository(db)
	router := http.NewServeMux()
	auth.NewAuthHandler(router, &auth.AuthHandlerDeps{
		Config: &cfg,
	})
	link.NewLinkHandler(router)
	server := http.Server{
		Addr:    cfg.WServer.Addr,
		Handler: router,
	}
	server.ListenAndServe()

}
