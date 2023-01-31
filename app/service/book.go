package service

import (
	"log"

	"github.com/nexters/book/app/entity"
	"github.com/nexters/book/app/repository"
	"github.com/nexters/book/external/search"
)

type (
	// BookService BookService Interface
	BookService interface {
		CreateBook(title string, ISBN string, userID string) (entity.Book, error)
		FindBookByISBN(ISBN string) (entity.Book, error)
		FindAllBooks(userID string, isReading bool) ([]entity.Book, error)
	}

	// bookService bookService Struct
	bookService struct {
		repo       repository.BookRepository
		bookSearch search.BookSearch
	}
)

// NewBookService 생성자
func NewBookService(repo repository.BookRepository, bs search.BookSearch) BookService {
	return bookService{repo, bs}
}

// CreateBook 책 추가
func (b bookService) CreateBook(title string, ISBN string, userID string) (entity.Book, error) {
	// search naver
	searchedRes, err := b.bookSearch.SearchBookByTitle(title)
	if err != nil {
		log.Fatal(err)
		return entity.Book{}, err
	}

	book := repository.CreateBookParams{
		UserID: userID,
	}

	// find one by matching ISBN
	for _, item := range searchedRes.Items {
		if item.ISBN == ISBN {
			book.Author = item.Author
			book.Description = item.Description
			book.ISBN = item.ISBN
			book.Title = item.Title
			book.Publisher = item.Publisher
			book.Price = item.Price
			book.ShopLink = item.ShopLink
			book.Image = item.Image

			break
		}
	}

	// add to db
	return b.repo.CreateBook(book)

	// TODO: add to user_books
}

// FindAllBooks 책 조회
func (b bookService) FindAllBooks(userID string, isReading bool) ([]entity.Book, error) {

	return b.repo.FindAllBooks(userID, isReading)
}

// FindBooksByISBN ISBN으로 책 조회
func (b bookService) FindBookByISBN(ISBN string) (entity.Book, error) {
	return b.repo.FindBookByISBN(ISBN)
}
