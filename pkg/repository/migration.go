package repository

import (
	"testAPI/models"

	"gorm.io/gorm"
)

func RunMigration(db *gorm.DB) error {
	return db.AutoMigrate(
		&models.Person{},
	)
}
