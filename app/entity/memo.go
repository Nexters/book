package entity

import "gorm.io/gorm"

type Memo struct {
	gorm.Model
	UserID   uint64 `gorm:"foreignKey" json:"userId"`
	BookID   uint64 `gorm:"foreignKey" json:"bookId"`
	Text     string `gorm:"text" json:"text"`
	Category string `gorm:"category" json:"category"`
}
