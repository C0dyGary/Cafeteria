package domain

import "gorm.io/gorm"

type OrderDetail struct {
	gorm.Model
	OrderID   uint    `json:"order_id"`
	Order     Order   `json:"order" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProductID uint    `json:"product_id"`
	Product   Product `json:"product" gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Quantity  int     `json:"quantity"`
	UnitPrice float64 `json:"unit_price"`
	Total     float64 `json:"total"`
}
