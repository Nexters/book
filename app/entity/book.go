package entity

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string `gorm:"title" json:"title"`
	Author      string `gorm:"author" json:"author"`
	ShopLink    string `gorm:"shopLink" json:"shopLink"`
	Image       string `gorm:"image" json:"image"`
	Price       string `gorm:"price" json:"price"`
	Publisher   string `gorm:"publisher" json:"publisher"`
	ISBN        string `gorm:"ISBN" json:"ISBN"`
	Description string `gorm:"description" json:"description"`
}
