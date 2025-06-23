package main

import (
	"fmt"
	"log"
	"os"
	"temp/internal/link"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main(){
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Файл .env не загружен: %s\n", err)
	}
	dsn := os.Getenv("DSN")
	db,err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		log.Fatalln("Error connect db: ",err)
	}
	db.AutoMigrate(&link.Link{})
}