package repository

import (
	"fmt"
	"goFinal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetAllUsers() (*[]model.Customer, error)
	GetUserByID(id int) (*model.Customer, error)
	GetUserByEmail(email string) (*model.Customer, error)
	InsertUser(data model.Customer) (int64, error)
	UpdateUserByAdmin(id int, dataUser map[string]interface{}) (int, error)
	UpdateUserByID(id int, dataUser model.Customer) (int, error)
	UpdateUserByEmail(email string, dataUser model.Customer) (int, error)
}
type countryDB struct {
	db *gorm.DB
}

func NewUsersRepository(gormdb *gorm.DB) UserRepository {
	return countryDB{db: gormdb}
}

// GetAll implements ConutryRepository.
func (connect countryDB) GetAllUsers() (*[]model.Customer, error) {
	user := []model.Customer{}
	result := connect.db.Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (connect countryDB) GetUserByEmail(email string) (*model.Customer, error) {
	user := model.Customer{}
	result := connect.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (connect countryDB) GetUserByID(id int) (*model.Customer, error) {
	user := model.Customer{}
	result := connect.db.Where("uid = ?", id).Find(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (connect countryDB) UpdateUserByID(id int, dataUser model.Customer) (int, error) {
	result := connect.db.Model(&model.Customer{}).Where("uid = ?", id).Updates(dataUser)
	if result.Error != nil {
		return -1, result.Error
	}
	return int(result.RowsAffected), nil
}

func (connect countryDB) UpdateUserByEmail(email string, dataUser model.Customer) (int, error) {
	fmt.Printf("email : %v", email)
	result := connect.db.Model(&model.Customer{}).Where("email = ?", email).Updates(dataUser)
	if result.Error != nil {
		return -1, result.Error
	}
	return int(result.RowsAffected), nil
}

func (connect countryDB) UpdateUserByAdmin(id int, dataUser map[string]interface{}) (int, error) {
	result := connect.db.Model(&model.Customer{}).Where("uid = ?", id).Updates(dataUser)
	if result.Error != nil {
		return -1, result.Error
	}
	return int(result.RowsAffected), nil
}

func (connect countryDB) InsertUser(data model.Customer) (int64, error) {
	user := data
	result := connect.db.Create(&user)
	if result.Error != nil {
		return -1, result.Error
	}
	return result.RowsAffected, nil

}
