package entity

import "gorm.io/gorm"

// Memo memo entity
type Memo struct {
	gorm.Model
	UserID   uint   `gorm:"foreignKey" json:"userId"`
	BookID   uint   `gorm:"foreignKey" json:"bookId"`
	Text     string `gorm:"column:text" json:"text"`
	Category string `gorm:"column:category" json:"category"`
}
