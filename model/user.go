package model

import (
	"time"
)

type User struct {
	ID           uint      `gorm:"column:uid;primaryKey;autoIncrement"`
	UserName     string    `gorm:"column:user_name;size:255;not null"`
	Email        string    `gorm:"column:email;size:255;not null"`
	Password     string    `gorm:"column:password;size:255;not null"`
	Gender       string    `gorm:"column:user_gender;size:10;not null"`
	Birthday     time.Time `gorm:"column:user_birthday;not null"`
	Image        string    `gorm:"column:user_img;size:1000;not null"`
	FavoriteFood string    `gorm:"column:favorite_food;size:500;not null"`
	UserType     int       `gorm:"column:user_type;not null"`
	GoogleID     string    `gorm:"column:googleid;not null"`
	Profile      string    `gorm:"column:user_profile;not null"`
	Cluster      int       `gorm:"column:user_cluster;not null"`
}

func (User) TableName() string {
	// Table name in database
	return "users"
}
