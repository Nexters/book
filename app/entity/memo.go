package entity

import "gorm.io/gorm"

type Memo struct {
	gorm.Model
	UserID   uint64 `gorm:"primaryKey" json:"userId"`
	BookID   uint64 `gorm:"primaryKey" json:"bookId"`
	Text     string `gorm:"text" json:"text"`
	Category string `gorm:"category" json:"category"`
}
