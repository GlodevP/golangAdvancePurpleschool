package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	WServer    WebserverConfig
	AuthConfig AuthConfig
	DBConfig DBConfig 
}

type AuthConfig struct {
	Secret string
}

type WebserverConfig struct {
	Addr string
}
type DBConfig struct{
	DSN string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Файл .env не загружен: %s\n", err)
	}
	return Config{
		WServer: WebserverConfig{
			Addr: os.Getenv("WebserverAddr"),
		},
		AuthConfig: AuthConfig{
			Secret: os.Getenv("Tocken"),
		},
		DBConfig: DBConfig{
			DSN: os.Getenv("DSN"),
		},
	}
}
