package entity

import "gorm.io/gorm"

// Memo memo entity
type Memo struct {
	gorm.Model
	BookID   uint   `gorm:"bookId"`
	Text     string `gorm:"column:text" json:"text"`
	Category string `gorm:"column:category" json:"category"`
}
