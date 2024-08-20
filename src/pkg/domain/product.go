package domain

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string        `json:"name"`
	Price       float64       `json:"price"`
	Description string        `json:"description"`
	Category    string        `json:"category"`
	OrderDetail []OrderDetail `json:"order_detail" gorm:"foreignKey:ProductID"`
}
