package order

import (
	"4-order-api/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
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

func (rep *Repository) Add(o *Product)error{
	err := rep.db.Create(o).Error
	if err != nil {
		return err
	}
	return nil
}

func (rep *Repository) GetByID(id uint)(*Product,error){
	var o Product
	err := rep.db.First(&o,"id = ?",id).Error
	if err != nil {
		return nil, err
	}
	return &o,nil
}


func (rep *Repository) Update(p *Product)(*Product,error){
	res := rep.db.Clauses(clause.Returning{}).Updates(p)
	if res.Error != nil {
		return nil, res.Error
	}
	return p, nil
}
func (rep *Repository) Delete(id uint)error{
	res := rep.db.Clauses(clause.Returning{}).Delete(&Product{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}