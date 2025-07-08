package main

import (
	"log"
	"net/http"
	"temp/configs"
	"temp/internal/auth"
	"temp/internal/link"
	"temp/pkg/db"
	"temp/pkg/middleware"
)

func main() {
	cfg := configs.LoadConfig()
	db, err := db.NewDB(cfg)
	if err != nil {
		log.Fatalln("Error, not connect to db,", err)
	}
	LinkRepository := link.NewLinkRepository(db)
	router := http.NewServeMux()
	auth.NewAuthHandler(router, &auth.AuthHandlerDeps{
		Config: &cfg,
	})
	link.NewLinkHandler(&link.LinkRepositoryDeps{
		Router:     router,
		Repository: LinkRepository,
	})
	stack := middleware.Chain(middleware.CORS, middleware.Logging)
	server := http.Server{
		Addr:    cfg.WServer.Addr,
		Handler: stack(router),
	}
	server.ListenAndServe()

}
