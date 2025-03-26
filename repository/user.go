package repository

import (
	"fmt"
	"goFinal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() (*[]model.Customer, error)
	GetUserByEmail(email string) (*model.Customer, error)
	InsertUser(data model.Customer) (int64, error)
	UpdatePasswordByEmail(email string, password string) (int, error)
}
type userDB struct {
	db *gorm.DB
}

func NewUsersRepository(gormdb *gorm.DB) UserRepository {
	return &userDB{db: gormdb}
}

// GetAll implements ConutryRepository.
func (connect userDB) GetAllUsers() (*[]model.Customer, error) {
	user := []model.Customer{}
	result := connect.db.Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (connect userDB) GetUserByEmail(email string) (*model.Customer, error) {
	user := model.Customer{}
	result := connect.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (connect userDB) UpdatePasswordByEmail(email string, password string) (int, error) {
	fmt.Printf("email : %v", email)
	result := connect.db.Model(&model.Customer{}).Where("email = ?", email).Update("password", password)
	fmt.Printf("result : %v", result)
	if result.Error != nil {
		return -1, result.Error
	}
	return int(result.RowsAffected), nil
}

func (connect userDB) InsertUser(data model.Customer) (int64, error) {
	user := data
	result := connect.db.Create(&user)
	if result.Error != nil {
		return -1, result.Error
	}
	return result.RowsAffected, nil

}
