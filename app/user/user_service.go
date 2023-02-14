package user

import (
	"time"

	"github.com/nexters/book/app/book"
)

type (
	UserService interface {
		FindUserStat(userID string) (UserStatPayload, error)
	}
	userService struct {
		userRepo    UserRepository
		bookService book.BookService
	}
)

// NewUserService 생성자
func NewUserService(u UserRepository, b book.BookService) UserService {
	return userService{u, b}
}

// FindUserStat 사용자의 통계 조회
func (u userService) FindUserStat(userID string) (stat UserStatPayload, err error) {
	user, err := u.userRepo.FindUserByUID(userID)
	if err != nil {
		return
	}
	start := user.CreatedAt

	books_reading, err := u.bookService.FindAllBooks(userID, true)
	if err != nil {
		return
	}

	books_read, err := u.bookService.FindAllBooks(userID, false)
	if err != nil {
		return
	}

	stat.ReadCount = books_read.Count
	var end time.Time
	for _, book := range books_read.Books {
		if end.Sub(book.UpdatedAt) < 0 {
			end = book.UpdatedAt
		}
	}

	stat.MemoCount = countAllMemos(books_reading, books_read)
	stat.Duration = int64(end.Sub(start).Hours() / 24)

	if end.Sub(start) < 0 {
		stat.Duration = int64(0)
		return
	}

	return
}

// countAllMemos 읽는 중, 완독 모두 카운트
func countAllMemos(readings book.FindAllBooksPayload, reads book.FindAllBooksPayload) (memoCount int) {
	for _, reading := range readings.Books {
		memoCount += reading.MemoCount
	}
	for _, read := range reads.Books {
		memoCount += read.MemoCount
	}

	return
}
