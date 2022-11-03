package migrations

import (
	"gorm.io/gorm"

	models "task3/internal/task3/models"
)

func Migrate(db *gorm.DB) {
	// Will automatically create missing columns
	db.AutoMigrate(&models.User{})
	// Can also be done explicidly
	db.Migrator().AddColumn(&models.User{}, "LastLogin")
}
