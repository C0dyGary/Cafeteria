package domain

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Name  string  `json:"name"`
	Phone string  `json:"phone"`
	Order []Order `json:"order" gorm:"foreignKey:ClientID"`
}
