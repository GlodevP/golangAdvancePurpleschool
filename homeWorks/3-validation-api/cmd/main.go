package main

import (
	"3-validation-api/configs"
	"3-validation-api/internal/verify"
	"net/http"
)

func main() {
	cfg := configs.NewConfigs()
	router := http.NewServeMux()
	verify.NewVerifyHandler(cfg, router)
	http.ListenAndServe(cfg.Url, router)
}
