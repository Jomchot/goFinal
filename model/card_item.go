package model

type CardItem struct {
	Cart_item_id uint   `gorm:"column:cart_item_id;primaryKey;autoIncrement"`
	Cart_id      uint   `gorm:"column:cart_id;size:255;not null"`
	Product_id   uint   `gorm:"column:product_id;size:255;not null"`
	Quantity     uint   `gorm:"column:quantity;size:255;not null"`
	Created_at   string `gorm:"column:created_at;size:20;not null"`
	Updated_at   string `gorm:"column:updated_at;size:20;not null"`
	// associations (foreign keys)
	Card    Card    `gorm:"foreignKey:Cart_id;references:ID;constraint:OnDelete:RESTRICT,OnUpdate:RESTRICT" json:"cart"`
	Product Product `gorm:"foreignKey:Product_id;references:ID;constraint:OnDelete:RESTRICT,OnUpdate:RESTRICT" json:"product"`
}

func (CardItem) TableName() string {
	return "cart_item"
}
