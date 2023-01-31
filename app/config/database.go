package config

import (
	"log"

	"gorm.io/gorm/logger"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

// Database db
type Database struct {
	*gorm.DB
}

// NewDatabase 생성자
func NewDatabase(settings *Settings, dialector SQLiteDialector) Database {
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err)
	}

	return Database{db}
}

// DBModule db 모듈
var DBModule = fx.Module("database",
	fx.Provide(NewDatabase, NewSQLiteDialector, NewMySQLDialector),
)
