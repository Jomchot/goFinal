package service

import (
	"goFinal/model"
	"goFinal/repository"

	"gorm.io/gorm"
)

type showCardService interface {
	GetProduct(description string, minPrice float64, maxPrice float64) (*[]model.Product, error)
}

func NewCardService(gormdb *gorm.DB) showCardService {
	return &showCard{db: gormdb}
}

type showCard struct {
	db *gorm.DB
}

func (c showCard) GetProduct(description string, minPrice float64, maxPrice float64) (*[]model.Product, error) {
	productRepo := repository.NewProductRepository(c.db)
	product, err := productRepo.GetProduct(description, minPrice, maxPrice)
	if err != nil {
		return nil, err
	}

	return product, nil
}
