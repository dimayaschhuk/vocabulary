package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {
	dsn := "host=localhost user=backend password=12345 dbname=backend port=5435"
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
