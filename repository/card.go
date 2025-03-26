package repository

import (
	"goFinal/model"

	"gorm.io/gorm"
)

type CardRepository interface {
	PostCardItem(data model.Customer) (int64, error)
}
type cardDB struct {
	db *gorm.DB
}

func NewCardRepository(gormdb *gorm.DB) CardRepository {
	return &cardDB{db: gormdb}
}

func (connect cardDB) PostCardItem(data model.Customer) (int64, error) {
	user := data
	result := connect.db.Create(&user)
	if result.Error != nil {
		return -1, result.Error
	}
	return result.RowsAffected, nil

}
