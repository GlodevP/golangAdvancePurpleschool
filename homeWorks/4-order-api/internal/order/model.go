package order

import (
	"gorm.io/gorm"
	"github.com/lib/pq"
)

type Product struct{
	gorm.Model
	Name string `json:"name" validate:"required" gorm:"not null"`
	Description string `json:"description"`
	Images pq.StringArray `gorm:"type:text[]" json:"images"`
}

func NewProduct(name string,description string)*Product{
	return &Product{
		Name: name,
		Description: description,
	}
}
