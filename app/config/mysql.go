package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type MySQLDialector gorm.Dialector

func NewMySQLDialector(settings *Settings) MySQLDialector {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%v)/%v?", settings.Database.User, settings.Database.Password, settings.Database.Port, settings.Database.Name)

	return mysql.Open(dsn)
}
