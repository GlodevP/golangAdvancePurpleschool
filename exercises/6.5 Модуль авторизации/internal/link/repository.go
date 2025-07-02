package link

import "temp/pkg/db"

type LinkRepository struct{
	db *db.DB
}

func NewLinkRepository(database *db.DB) *LinkRepository{
	return &LinkRepository{db: database}
}

func (rep *LinkRepository) Create(link *Link){

}