package routes

import "gorm.io/gorm"

type Router struct {
	BaseData *gorm.DB
}
