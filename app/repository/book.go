package repository

import (
	"errors"

	"github.com/nexters/book/app/config"
	"github.com/nexters/book/app/entity"
)

// TODO: implement
type (
	// BookRepository BookRepository Interface
	BookRepository interface {
		CreateBook(params CreateBookParams) (entity.Book, error)
		FindAllBooks(userID string, isReading bool) ([]entity.Book, error)
		FindBookByISBN(ISBN string) (entity.Book, error)
		FindBookAndAllMemosByBookID(bookID uint) (entity.Book, error)
	}

	// bookRepository bookRepository Struct
	bookRepository struct {
		db config.Database
	}
)

// CreateBookParams 책 생성 parameters
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

// NewBookRepository 생성자
func NewBookRepository(db config.Database) BookRepository {
	return bookRepository{db}
}

// CreateBook 책 생성
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

	user := entity.User{}
	tx := b.db.Where("uid = ?", params.UserID).First(&user)

	if tx.Error != nil || tx.RowsAffected == 0 {
		err = errors.New("User not found")
		return
	}

	book.UserID = user.ID
	tx = b.db.Create(&book)

	return
}

// FindAllBooks 책 조회
func (b bookRepository) FindAllBooks(userID string, isReading bool) (books []entity.Book, err error) {
	user := entity.User{}
	res := b.db.Where("uid = ?", userID).First(&user)
	books = []entity.Book{}

	if res.Error != nil || res.RowsAffected == 0 {
		err = errors.New("User not found")
		return
	}

	err = b.db.Model(&user).Where("user_id = ? AND is_reading = ?", user.ID, isReading).Association("Books").Find(&books)
	if err != nil {
		err = errors.New("Books not found")
		return
	}
	return
}

// FindBookAndAllMemosByBookID bookID로 유저의 책과 모든 메모 조회
func (b bookRepository) FindBookAndAllMemosByBookID(bookID uint) (book entity.Book, err error) {
	tx := b.db.Preload("Memos").Where("books.id", bookID).First(&book)
	if tx.Error != nil {
		err = tx.Error
		return
	}

	return
}

// FindBookByISBN ISBN으로 책 조회
func (b bookRepository) FindBookByISBN(ISBN string) (book entity.Book, err error) {
	book = entity.Book{}
	tx := b.db.Where("isbn = ?", ISBN).First(&book)

	if tx.Error != nil {
		err = tx.Error
		return
	}

	if tx.RowsAffected == 0 {
		err = errors.New("Book not found")
		return
	}
	return
}
