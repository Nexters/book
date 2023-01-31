package service

import (
	"log"

	"github.com/nexters/book/app/entity"
	"github.com/nexters/book/app/repository"
	"github.com/nexters/book/external/search"
)

type (
	BookService interface {
		CreateBook(title string, ISBN string, userID string) (entity.Book, error)
		FindBookByISBN(ISBN string) (entity.Book, error)
		FindAllBooks(userID string) ([]entity.Book, error)
	}
	bookService struct {
		repo       repository.BookRepository
		bookSearch search.BookSearch
	}
)

func NewBookService(repo repository.BookRepository, bs search.BookSearch) BookService {
	return bookService{repo, bs}
}

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

func (b bookService) FindAllBooks(userID string) ([]entity.Book, error) {
	return b.repo.FindAllBooks(userID)
}

func (b bookService) FindBookByISBN(ISBN string) (entity.Book, error) {
	return b.repo.FindBookByISBN(ISBN)
}
