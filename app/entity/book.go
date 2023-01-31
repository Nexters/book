package entity

import "gorm.io/gorm"

// Book book entity
type Book struct {
	gorm.Model
	Title       string `gorm:"column:title" json:"title"`
	Author      string `gorm:"column:author" json:"author"`
	ShopLink    string `gorm:"column:shopLink" json:"shopLink"`
	Image       string `gorm:"column:image" json:"image"`
	Price       string `gorm:"column:price" json:"price"`
	Publisher   string `gorm:"column:publisher" json:"publisher"`
	ISBN        string `gorm:"column:ISBN" json:"ISBN"`
	Description string `gorm:"column:description" json:"description"`
}
