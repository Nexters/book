package config

import (
	"log"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

type (
	Database interface {
		MakeMigration(repositories ...interface{})
	}
	database struct {
		gorm *gorm.DB
	}
)

func NewDatabase(settings *Settings, dialector SQLiteDialector) Database {
	db, err := gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return database{db}
}

func (db database) MakeMigration(repositories ...interface{}) {
	if err := db.gorm.AutoMigrate(repositories...); err != nil {
		log.Fatal(err)
	}
}

var DBModule = fx.Module("database",
	fx.Provide(NewDatabase, NewSQLiteDialector),
)
