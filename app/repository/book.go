package repository

import (
	"errors"

	"github.com/nexters/book/app/config"
	"github.com/nexters/book/app/entity"
)

// TODO: implement
type (
	BookRepository interface {
		CreateBook(params CreateBookParams) (entity.Book, error)
		FindAllBooks(userID string) ([]entity.Book, error)
		FindBookByISBN(ISBN string) (entity.Book, error)
	}
	bookRepository struct {
		db config.Database
	}
)

type CreateBookParams struct {
	UserID      string `json:"userId"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	ShopLink    string `json:"shopLink"`
	Image       string `json:"image"`
	Price       string `json:"price"`
	Publisher   string `json:"publisher"`
	ISBN        string `json:"ISBN"`
	Description string `json:"description"`
}

func NewBookRepository(db config.Database) BookRepository {
	return bookRepository{db}
}

func (b bookRepository) CreateBook(params CreateBookParams) (entity.Book, error) {
	book := entity.Book{
		Title:       params.Title,
		Author:      params.Author,
		ShopLink:    params.ShopLink,
		Image:       params.Image,
		Price:       params.Price,
		Publisher:   params.Publisher,
		ISBN:        params.ISBN,
		Description: params.Description,
	}

	res := b.db.Where("isbn = ?", book.ISBN).FirstOrCreate(&book)

	if res.Error != nil || res.RowsAffected == 0 {
		return book, errors.New("Book creation failed")
	}

	user := entity.User{}
	res = b.db.Where("uid = ?", params.UserID).First(&user)

	if res.Error != nil || res.RowsAffected == 0 {
		return book, errors.New("User not found")
	}

	userBooks := entity.UserBooks{
		UserID: uint64(user.ID),
		BookID: uint64(book.ID),
	}

	res = b.db.Create(&userBooks)
	if res.Error != nil || res.RowsAffected == 0 {
		return book, errors.New("Book creation failed")
	}

	return book, nil
}

func (b bookRepository) FindAllBooks(userID string) ([]entity.Book, error) {
	user := entity.User{}
	res := b.db.Where("uid = ?", userID).First(&user)
	books := []entity.Book{}

	if res.Error != nil || res.RowsAffected == 0 {
		return books, errors.New("User not found")
	}

	err := b.db.Model(&user).Where("user_id = ?", user.ID).Association("Books").Find(&books)
	if err != nil {
		return books, errors.New("Books not found")
	}

	return books, nil
}

func (b bookRepository) FindBookByISBN(ISBN string) (entity.Book, error) {
	book := entity.Book{}
	res := b.db.Where("isbn = ?", ISBN).First(&book)

	if res.Error != nil {
		return book, res.Error
	}

	if res.RowsAffected == 0 {
		return book, errors.New("Book not found")
	}

	return book, nil
}
