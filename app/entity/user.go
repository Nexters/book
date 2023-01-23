package entity

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email string `gorm:"email" json:"email"`
	/*
		TODO: 카카오 인증, 애플 인증
	*/
	Books Book `gorm:"many2many:user_books" json:"books"`
}
