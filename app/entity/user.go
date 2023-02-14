package entity

import (
	"gorm.io/gorm"
)

// User user entity
type User struct {
	gorm.Model
	Email string `gorm:"column:email" json:"email"`
	Uid   string `gorm:"column:uid" json:"token"`
	Books []Book `json:"books"`
}
