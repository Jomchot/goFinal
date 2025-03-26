package repository

import (
	"goFinal/model"

	"gorm.io/gorm"
)

type ProductRepository interface {
	GetProduct(description string, minPrice, maxPrice float64) (*[]model.Product, error)
	// GetUserByEmail(email string) (*model.Customer, error)
	// InsertUser(data model.Customer) (int64, error)
	// UpdatePasswordByEmail(email string, password string) (int, error)
}
type productDB struct {
	db *gorm.DB
}

func NewProductRepository(gormdb *gorm.DB) ProductRepository {
	return &productDB{db: gormdb}
}

func (connect *productDB) GetProduct(description string, minPrice float64, maxPrice float64) (*[]model.Product, error) {
	products := []model.Product{}
	result := connect.db.Where("description LIKE ? AND price BETWEEN ? AND ?", "%"+description+"%", minPrice, maxPrice).Find(&products)
	if result.Error != nil {
		return nil, result.Error
	}
	return &products, nil
}
