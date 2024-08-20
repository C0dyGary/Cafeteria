package orm

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("cafeteria.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func Migrate(list []interface{}, db *gorm.DB) {
	err := db.AutoMigrate(list...)
	if err != nil {
		panic("failed to migrate database")
	}
}
