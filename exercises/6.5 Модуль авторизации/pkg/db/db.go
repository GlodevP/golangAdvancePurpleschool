package db

import (
	"temp/configs"

	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)

type DB struct{
	*gorm.DB
}

func NewDB(conf configs.Config) (*DB,error) {
	db,err := gorm.Open(postgres.Open(conf.DBConfig.DSN))
	if err != nil {
		return nil, err
	}
	return &DB{db},nil
}