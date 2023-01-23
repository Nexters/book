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
		FindAllBooks()
		FindBookByID(id uint64) (entity.Book, error)
	}
	bookRepository struct {
		db config.Database
	}
)

type CreateBookParams struct {
	Title       string `json:"title"`
	Author      string `json:"author"`
	ShopLink    string `json:"shopLink"`
	Image       string `json:"image"`
	Price       int    `json:"price"`
	Publisher   string `json:"publisher"`
	ISBN        string `json:"ISBN"`
	Description string `json:"discription"`
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

	res := b.db.Create(&book)

	if res.Error != nil || res.RowsAffected == 0 {
		return book, errors.New("Book creation failed")
	}

	return book, nil
}

func (b bookRepository) FindAllBooks() {}

func (b bookRepository) FindBookByID(id uint64) (entity.Book, error) {
	book := entity.Book{}
	res := b.db.Where("id = ?", id).First(&book)

	if res.Error != nil {
		return book, res.Error
	}

	if res.RowsAffected == 0 {
		return book, errors.New("Book not found")
	}

	return book, nil
}
