package config

import (
	"log"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

// Database db
type Database struct {
	*gorm.DB
}

// NewDatabase 생성자
func NewDatabase(settings *Settings, dialector SQLiteDialector) Database {
	db, err := gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return Database{db}
}

// DBModule db 모듈
var DBModule = fx.Module("database",
	fx.Provide(NewDatabase, NewSQLiteDialector, NewMySQLDialector),
)
