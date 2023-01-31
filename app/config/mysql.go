package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MySQLDialector MySQL Dialector
type MySQLDialector gorm.Dialector

// NewMySQLDialector 생성자
func NewMySQLDialector(settings *Settings) MySQLDialector {
	dsn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:%v)/%v?", settings.Database.User, settings.Database.Password, settings.Database.Port, settings.Database.Name)

	return mysql.Open(dsn)
}
