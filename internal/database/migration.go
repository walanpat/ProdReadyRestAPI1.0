package database

import (
	"ProdReadyRestAPI1.0/internal/comment"
	"github.com/jinzhu/gorm"
)

// MigrateDB - migrates our database and creates our comment table
func MigrateDB(db *gorm.DB) error {
	if result := db.AutoMigrate(&comment.Comment{}); result.Error != nil {
		return result.Error
	}
	return nil
}
