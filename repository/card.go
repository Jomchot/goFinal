package repository

import (
	"goFinal/model"

	"gorm.io/gorm"
)

type CardRepository interface {
	GetProductUsers() (*[]model.Customer, error)
	// GetUserByEmail(email string) (*model.Customer, error)
	// InsertUser(data model.Customer) (int64, error)
	// UpdatePasswordByEmail(email string, password string) (int, error)
}
type cardDB struct {
	db *gorm.DB
}

func NewCardRepository(gormdb *gorm.DB) CardRepository {
	return &cardDB{db: gormdb}
}

// GetAll implements ConutryRepository.
func (connect cardDB) GetProductUsers() (*[]model.Customer, error) {
	user := []model.Customer{}
	result := connect.db.Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
