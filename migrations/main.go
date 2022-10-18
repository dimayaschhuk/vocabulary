package migrations

import "gorm.io/gorm"

type Word struct {
	Id        int
	Name      string
	Translate string
	Lesson    int
}

func Create(db *gorm.DB) error {
	return db.AutoMigrate(&Word{})
}
