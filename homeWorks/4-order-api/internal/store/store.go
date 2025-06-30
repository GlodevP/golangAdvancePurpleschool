package store

import (
	"4-order-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct{
	db *gorm.DB
}

func NewDB(cfg config.Config) (*DB,error){
	db,err := gorm.Open(postgres.Open(cfg.DB.DSN))
	if err != nil {
		return nil,err
	}
	return &DB{
		db: db,
	},nil
}

func (s *DB) GetProduct(id string)(*Product,error){
	var o Product
	err := s.db.First(&o,"id = ?",id).Error
	if err != nil {
		return nil, err
	}
	return &o,nil
}
func (s *DB) AddProduct(o *Product)error{
	err := s.db.Create(o).Error
	if err != nil {
		return err
	}
	return nil
}