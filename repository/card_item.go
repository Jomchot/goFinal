package repository

import (
	"goFinal/model"

	"gorm.io/gorm"
)

type CardItemRepository interface {
	GetCardItem() (*[]model.Customer, error)
	// GetUserByEmail(email string) (*model.Customer, error)
	// InsertUser(data model.Customer) (int64, error)
	// UpdatePasswordByEmail(email string, password string) (int, error)
}
type cardItemDB struct {
	db *gorm.DB
}

func NewCardItemRepository(gormdb *gorm.DB) CardItemRepository {
	return &cardItemDB{db: gormdb}
}

// GetAll implements ConutryRepository.
func (connect cardItemDB) GetCardItem() (*[]model.Customer, error) {
	user := []model.Customer{}
	result := connect.db.Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
