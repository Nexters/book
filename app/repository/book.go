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

func (b bookRepository) CreateBook(params CreateBookParams) (book entity.Book, err error) {
	book = entity.Book{
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
		err = errors.New("Book creation failed")
		return
	}

	user := entity.User{}
	res = b.db.Where("uid = ?", params.UserID).First(&user)

	if res.Error != nil || res.RowsAffected == 0 {
		err = errors.New("User not found")
		return
	}

	userBooks := entity.UserBooks{
		UserID: uint64(user.ID),
		BookID: uint64(book.ID),
	}

	res = b.db.Create(&userBooks)
	if res.Error != nil || res.RowsAffected == 0 {
		err = errors.New("Book creation failed")
		return
	}
	return
}

func (b bookRepository) FindAllBooks(userID string) (books []entity.Book, err error) {
	user := entity.User{}
	res := b.db.Where("uid = ?", userID).First(&user)
	books = []entity.Book{}

	if res.Error != nil || res.RowsAffected == 0 {
		err = errors.New("User not found")
		return
	}

	err = b.db.Model(&user).Where("user_id = ?", user.ID).Association("Books").Find(&books)
	if err != nil {
		err = errors.New("Books not found")
		return
	}
	return
}

func (b bookRepository) FindBookByISBN(ISBN string) (book entity.Book, err error) {
	book = entity.Book{}
	res := b.db.Where("isbn = ?", ISBN).First(&book)

	if res.Error != nil {
		err = res.Error
		return
	}

	if res.RowsAffected == 0 {
		err = errors.New("Book not found")
		return
	}
	return
}
