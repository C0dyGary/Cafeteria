package domain

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	ClientID    uint          `json:"client_id"`
	Client      Client        `json:"client" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	OrderDetail []OrderDetail `json:"order_detail" gorm:"foreignKey:OrderID"`
	Total       float64       `json:"total"`
}
