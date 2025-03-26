package repository

import (
	"errors"
	"goFinal/model"
	"time"

	"gorm.io/gorm"
)

type CardItemRepository interface {
	PostCardItem(cartID uint, item uint, quantity uint) (*model.CardItem, error)
}
type cardItemDB struct {
	db *gorm.DB
}

func NewCardItemRepository(gormdb *gorm.DB) CardItemRepository {
	return &cardItemDB{db: gormdb}
}

func (connect cardItemDB) PostCardItem(cartID uint, item uint, quantity uint) (*model.CardItem, error) {
	// ตรวจสอบว่ามีสินค้านี้ใน Cart อยู่แล้วหรือไม่
	var product model.CardItem
	err := connect.db.Where("cart_id = ? AND product_id = ?", cartID, item).First(&product).Error
	if err == nil {
		return nil, errors.New("มีสินค้านี้อยู่ในตระกร้าแล้ว")
	}

	// ถ้าไม่พบสินค้าก็สร้าง item ใหม่
	itemCard := model.CardItem{
		Cart_id:    cartID,
		Product_id: item,
		Quantity:   quantity,
		Created_at: time.Now().Format("2006-01-02 15:04:05"),
		Updated_at: time.Now().Format("2006-01-02 15:04:05"),
	}

	// บันทึก item ใหม่
	result := connect.db.Create(&itemCard)
	if result.Error != nil {
		return nil, result.Error
	}

	return &itemCard, nil
}
