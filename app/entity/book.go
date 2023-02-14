package entity

import (
	"gorm.io/gorm"
)

// Book book entity
type Book struct {
	gorm.Model
	UserID      uint   `json:"userId"`
	Title       string `gorm:"column:title" json:"title"`
	Author      string `gorm:"column:author" json:"author"`
	ShopLink    string `gorm:"column:shopLink" json:"shopLink"`
	Image       string `gorm:"column:image" json:"image"`
	Price       string `gorm:"column:price" json:"price"`
	Publisher   string `gorm:"column:publisher" json:"publisher"`
	ISBN        string `gorm:"column:ISBN" json:"ISBN"`
	Description string `gorm:"column:description" json:"description"`
	IsReading   bool   `gorm:"column:is_reading;default:true" json:"isReading"`
	Memos       []Memo `json:"memos"`
}
