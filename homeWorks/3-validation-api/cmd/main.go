package main

import (
	"3-validation-api/configs"
	"3-validation-api/internal/verify"
	"log"
	"net/http"
)

func main() {
	cfg := configs.NewConfigs()
	router := http.NewServeMux()
	err := verify.NewVerifyHandler(cfg, router)
	if err != nil {
		log.Fatalln(err)
	}
	http.ListenAndServe(cfg.Url, router)
}
