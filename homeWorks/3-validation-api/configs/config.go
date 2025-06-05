package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct{
	Url string
	EmailConfig EmailConfig
}

type EmailConfig struct{
	EmailServer string
	EmailSendler string
	EmailSecret string
}

func NewConfigs() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Файл .env не загружен: %s\n", err)
	}
	return &Config{
		Url: os.Getenv("WEB_SERVER_URL"),
		EmailConfig: EmailConfig{
			EmailServer: os.Getenv("MAIL_SERVER"),
			EmailSendler: os.Getenv("EMAIL_SENDLER"),
			EmailSecret: os.Getenv("EMAIL_SENCRET"),
		},
	}
}