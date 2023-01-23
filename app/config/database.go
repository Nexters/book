package config

import (
	"log"

	"go.uber.org/fx"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func NewDatabase(settings *Settings, dialector SQLiteDialector) Database {
	db, err := gorm.Open(dialector, &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return Database{db}
}

var DBModule = fx.Module("database",
	fx.Provide(NewDatabase, NewSQLiteDialector, NewMySQLDialector),
)
