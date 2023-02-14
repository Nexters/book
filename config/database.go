package config

import (
	"log"

	"github.com/nexters/book/config/environment"
	"gorm.io/gorm/logger"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

// Database db
type Database struct {
	*gorm.DB
}

// NewDatabase 생성자
func NewDatabase(settings *Settings, dialector MySQLDialector) Database {
	config := gorm.Config{}
	if settings.App.ENV != environment.PROD {
		config.Logger = logger.Default.LogMode(logger.Info)
	}
	db, err := gorm.Open(dialector, &config)

	if err != nil {
		log.Fatal(err)
	}

	return Database{db}
}

// DBModule db 모듈
var DBModule = fx.Module("database",
	fx.Provide(NewDatabase, NewSQLiteDialector, NewMySQLDialector),
)
