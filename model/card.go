package model

import "time"

type Card struct {
	Cart_id     uint      `gorm:"column:cart_id;primaryKey;autoIncrement"`
	Customer_id uint      `gorm:"column:customer_id;size:255;not null"`
	Cart_name   string    `gorm:"column:cart_name;size:255;not null"`
	Created_at  time.Time `gorm:"column:created_at;size:20;not null;autoUpdateTime"`
	Updated_at  time.Time `gorm:"column:updated_at;size:20;not null;autoUpdateTime"`

	Customer Customer `gorm:"foreignKey:customer_id;references:customer_id;constraint:OnDelete:RESTRICT,OnUpdate:RESTRICT  " json:"customer"`
}

func (Card) TableName() string {
	return "cart"
}
