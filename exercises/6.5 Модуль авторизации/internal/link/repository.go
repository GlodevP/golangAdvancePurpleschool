package link

import (
	"gorm.io/gorm/clause"
	"temp/pkg/db"
)

type LinkRepository struct {
	db *db.DB
}

func NewLinkRepository(database *db.DB) *LinkRepository {
	return &LinkRepository{db: database}
}

func (rep *LinkRepository) Create(link *Link) (*Link, error) {
	res := rep.db.Create(link)
	if res.Error != nil {
		return nil, res.Error
	}
	return link, nil
}

func (rep *LinkRepository) GetByHash(hash string) (*Link, error) {
	var link Link
	res := rep.db.First(&link, "hash = ?", hash)
	if res.Error != nil {
		return nil, res.Error
	}
	return &link, nil
}
func (rep *LinkRepository) GetByID(id uint) (*Link, error) {
	var link Link
	res := rep.db.First(&link, id)
	if res.Error != nil {
		return nil, res.Error
	}
	return &link, nil
}
func (rep *LinkRepository) Update(link *Link) (*Link, error) {
	res := rep.db.Clauses(clause.Returning{}).Updates(link)
	if res.Error != nil {
		return nil, res.Error
	}
	return link, nil
}

func (rep *LinkRepository) Delete(id uint) error {
	res := rep.db.Clauses(clause.Returning{}).Delete(&Link{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}
