package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string `gorm:"email" json:"email"`
	Uid   string `gorm:"uid" json:"uid"`
	Books []Book `gorm:"many2many:user_books;" json:"books"`
}