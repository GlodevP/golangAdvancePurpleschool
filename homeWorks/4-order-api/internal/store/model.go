package store

import "gorm.io/gorm"

type Product struct{
	gorm.Model
	Name string `json:"name" validate:"required" gorm:"not null"`
	Description string `json:"description"`
}

func NewOrder(name string,description string)*Product{
	return &Product{
		Name: name,
		Description: description,
	}
}
