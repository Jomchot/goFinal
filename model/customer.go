package model

import (
	"time"
)

type Customer struct {
	CustomerID  uint      `gorm:"column:customer_id;primaryKey;autoIncrement"`
	FirstName   string    `gorm:"column:first_name;size:255;not null"`
	LastName    string    `gorm:"column:last_name;size:255;not null"`
	Email       string    `gorm:"column:email;size:255;not null"`
	PhoneNumber string    `gorm:"column:phone_number;size:20;not null"`
	Address     string    `gorm:"column:address;size:255;not null"`
	Password    string    `gorm:"column:password;size:255;not null"`
	CreatedAt   time.Time `gorm:"column:created_at;not null;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;not null;autoUpdateTime"`
}

func (Customer) TableName() string {
	// Table name in database
	return "customer"
}
