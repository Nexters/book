package config

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// SQLiteDialector sqlite dialector
type SQLiteDialector gorm.Dialector

// NewSQLiteDialector 생성자
func NewSQLiteDialector(settings *Settings) SQLiteDialector {
	db := fmt.Sprintf("%s.sqlite", settings.Database.Name)
	return sqlite.Open(db)
}
