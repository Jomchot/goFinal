package model

import "time"

type Product struct {
	Product_id     uint      `gorm:"column:product_id;primaryKey;autoIncrement"`
	Product_name   string    `gorm:"column:description;size:255;not null"`
	Description    string    `gorm:"column:description;size:255;not null"`
	Price          float32   `gorm:"column:price;size:255;not null"`
	Stock_quantity string    `gorm:"column:stock_quantity;size:20;not null"`
	Created_at     time.Time `gorm:"column:created_at;not null;autoCreateTime"`
	Updated_at     time.Time `gorm:"column:updated_at;not null;autoUpdateTime"`
}

func (Product) TableName() string {
	// Table name in database
	return "product"
}
