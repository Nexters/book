package book

import (
	"github.com/nexters/book/app/entity"
	"github.com/nexters/book/external/search"
)

type (
	// BookService BookService Interface
	BookService interface {
		CreateBook(ISBN string, userID string) (entity.Book, error)
		FindBookByISBN(ISBN string, category string) (FindBookPayload, error)
		FindAllBooks(userID string, isReading bool) (FindAllBooksPayload, error)
		FindBookAndAllMemosByBookID(bookID uint, category string) (FindBookPayload, error)
		UpdateBook(bookID uint, isReading bool) (entity.Book, error)
		DeleteBook(bookID uint, userID string) (entity.Book, error)
	}

	// bookService bookService Struct
	bookService struct {
		repo       BookRepository
		bookSearch search.BookSearch
	}
)

// NewBookService 생성자
func NewBookService(repo BookRepository, bs search.BookSearch) BookService {
	return bookService{repo, bs}
}

// CreateBook 책 추가
func (b bookService) CreateBook(ISBN string, userID string) (entity.Book, error) {
	// search naver
	searchedRes, err := b.bookSearch.SearchBook(ISBN)
	if err != nil {
		return entity.Book{}, err
	}

	book := CreateBookParams{
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
func (b bookService) FindAllBooks(userID string, isReading bool) (payload FindAllBooksPayload, err error) {

	books, err := b.repo.FindAllBooks(userID, isReading)
	if err != nil {
		return
	}

	payload.Books = make([]FindBookPayload, 0)
	for _, book := range books {
		Memocount := len(book.Memos)
		bookPayload := FindBookPayload{
			Book:      book,
			MemoCount: Memocount,
		}

		payload.Books = append(payload.Books, bookPayload)
		payload.Count++
	}

	return
}

// FindBookAndAllMemosByBookID
func (b bookService) FindBookAndAllMemosByBookID(bookID uint, category string) (payload FindBookPayload, err error) {
	book, err := b.repo.FindBookAndAllMemosByBookID(bookID)
	if err != nil {
		return
	}

	memoCount := len(book.Memos)
	payload = FindBookPayload{
		Book:      book,
		MemoCount: memoCount,
	}

	// if not category, return all memos
	if category == "" {
		return
	}

	// if category, filter category
	memos := make([]entity.Memo, 0)
	for _, memo := range book.Memos {
		if memo.Category == category {
			memos = append(memos, memo)
		}
	}

	payload.Memos = memos
	return
}

// FindBooksByISBN ISBN으로 책 조회
func (b bookService) FindBookByISBN(ISBN string, category string) (payload FindBookPayload, err error) {
	book, err := b.repo.FindBookByISBN(ISBN)
	if err != nil {
		return
	}

	memoCount := len(book.Memos)
	payload = FindBookPayload{
		Book:      book,
		MemoCount: memoCount,
	}

	// if not category, return all memos
	if category == "" {
		return
	}

	// if category, filter category
	memos := make([]entity.Memo, 0)
	for _, memo := range book.Memos {
		if memo.Category == category {
			memos = append(memos, memo)
		}
	}

	payload.Memos = memos
	return
}

// UpdateBook 책의 읽는 중/완독 업데이트
func (b bookService) UpdateBook(bookID uint, isReading bool) (book entity.Book, err error) {
	book, err = b.repo.UpdateBook(bookID, isReading)

	return
}

// DeleteBook 책 삭제
func (b bookService) DeleteBook(bookID uint, userID string) (entity.Book, error) {
	return b.repo.DeleteBook(bookID, userID)
}
