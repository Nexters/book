package entity

type Book struct {
	BookID      uint64 `gorm:"id" json:"bookId"`
	Title       string `gorm:"title" json:"title"`
	Author      string `gorm:"author" json:"author"`
	ShopLink    string `gorm:"shopLink" json:"shopLink"`
	Image       string `gorm:"image" json:"image"`
	Price       int    `gorm:"price" json:"price"`
	Publisher   string `gorm:"publisher" json:"publisher"`
	ISBN        string `gorm:"ISBN" json:"ISBN"`
	Description string `gorm:"description" json:"discription"`
}
