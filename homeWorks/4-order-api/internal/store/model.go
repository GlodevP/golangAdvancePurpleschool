package store

import "gorm.io/gorm"

type Order struct{
	gorm.Model
	Name string `json:"name" validate:"required" gorm:"not null"`
	Description string `json:"description"`
}

func NewOrder(name string,description string)*Order{
	return &Order{
		Name: name,
		Description: description,
	}
}
