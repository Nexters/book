package config

import (
	"log"

	"gorm.io/driver/sqlite"
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

func NewDatabase() database {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

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
