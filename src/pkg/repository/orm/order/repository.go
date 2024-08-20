package order

import (
	"Cafeteria/src/pkg/domain"
	"gorm.io/gorm"
)

type Repository struct {
	DB *gorm.DB
}

func (r Repository) GetClientByID(id uint, d *domain.Client) error {
	//TODO implement me
	panic("implement me")
}
