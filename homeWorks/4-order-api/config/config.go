package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct{
	Webserver *WebserverConfig
	DB *DBConfig
}

type WebserverConfig struct{
	Addr string
}
type DBConfig struct{
	DSN string
}

func NewConfigs() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Файл .env не загружен: %s\n", err)
	}
	return &Config{
		Webserver: &WebserverConfig{
			Addr: os.Getenv("WEB_SERVER_URL"),
		},
		DB: &DBConfig{
			DSN: os.Getenv("DSN"),
		},
		}
}