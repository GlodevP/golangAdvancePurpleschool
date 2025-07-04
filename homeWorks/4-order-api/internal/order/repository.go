package order

import (
	"4-order-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Repository struct{
	db *gorm.DB
}

func NewRepository(cfg config.Config) (*Repository,error){
	db,err := gorm.Open(postgres.Open(cfg.DB.DSN))
	if err != nil {
		return nil,err
	}
	return &Repository{
		db: db,
	},nil
}

func (s *Repository) GetProduct(id string)(*Product,error){
	var o Product
	err := s.db.First(&o,"id = ?",id).Error
	if err != nil {
		return nil, err
	}
	return &o,nil
}
func (s *Repository) AddProduct(o *Product)error{
	err := s.db.Create(o).Error
	if err != nil {
		return err
	}
	return nil
}