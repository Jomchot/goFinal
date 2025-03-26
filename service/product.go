package service

import (
	"goFinal/model"
	"goFinal/repository"

	"gorm.io/gorm"
)

type showProductService interface {
	GetProduct(description string, minPrice float64, maxPrice float64) (*[]model.Product, error)
}

func NewProductService(gormdb *gorm.DB) showDataService {
	return showData{db: gormdb}
}

type showProduct struct {
	db *gorm.DB
}

func (c showData) GetProduct(description string, minPrice float64, maxPrice float64) (*[]model.Product, error) {
	productRepo := repository.NewProductRepository(c.db)
	product, err := productRepo.GetProduct(description, minPrice, maxPrice)
	if err != nil {
		return nil, err
	}

	return product, nil
}
