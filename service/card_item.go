package service

import (
	"goFinal/model"
	"goFinal/repository"

	"gorm.io/gorm"
)

type showCardItemService interface {
	PostCardItem(cartID uint, productID uint, quantity uint) (*model.CardItem, error)
}

func NewCardItemService(gormdb *gorm.DB) showCardItemService {
	return &showCardItem{db: gormdb}
}

type showCardItem struct {
	db *gorm.DB
}

// func (c showCard) GetProduct(description string, minPrice float64, maxPrice float64) (*[]model.Product, error) {
// 	productRepo := repository.NewProductRepository(c.db)
// 	product, err := productRepo.GetProduct(description, minPrice, maxPrice)
// 	if err != nil {
// 		return nil, err
// 	}

//		return product, nil
//	}
func (c *showCardItem) PostCardItem(cartID uint, productID uint, quantity uint) (*model.CardItem, error) {
	cardItemRepo := repository.NewCardItemRepository(c.db)
	// Call the repository to save the card item
	item, err := cardItemRepo.PostCardItem(cartID, productID, quantity)
	if err != nil {
		return nil, err
	}

	return item, nil
}
