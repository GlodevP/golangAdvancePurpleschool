package store

import (
	"4-order-api/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct{
	db *gorm.DB
}

func NewDB(cfg config.Config) *DB{
	db,err := gorm.Open(postgres.Open(cfg.DB.DSN))
	if err != nil {
		log.Fatalln("Error connect db: ",err)
	}
	return &DB{
		db: db,
	}
}

func (s *DB) GetOrder(id string)(*Order,error){
	var o Order
	err := s.db.First(&o,"id = ?",id).Error
	if err != nil {
		return nil, err
	}
	return &o,nil
}
func (s *DB) AddOrder(o *Order)error{
	err := s.db.Create(o).Error
	if err != nil {
		return err
	}
	return nil
}