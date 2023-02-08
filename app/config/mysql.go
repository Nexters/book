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
	dsn := fmt.Sprintf("%s:%s@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", settings.Database.User, settings.Database.Password, settings.Database.URL, settings.Database.Port, settings.Database.Name)

	return mysql.Open(dsn)
}
