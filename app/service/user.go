package service

import (
	"time"

	"github.com/nexters/book/app/repository"
	"github.com/nexters/book/app/service/payloads"
)

type (
	UserService interface {
		FindUserStat(userID string) (payloads.UserStatPayload, error)
	}
	userService struct {
		userRepo    repository.UserRepository
		bookService BookService
	}
)

func NewUserService(u repository.UserRepository, b BookService) UserService {
	return userService{u, b}
}

func (u userService) FindUserStat(userID string) (stat payloads.UserStatPayload, err error) {
	user, err := u.userRepo.FindUserByUID(userID)
	if err != nil {
		return
	}
	start := user.CreatedAt

	books, err := u.bookService.FindAllBooks(userID, false)
	if err != nil {
		return
	}

	stat.ReadCount = books.Count
	var end time.Time
	for _, book := range books.Books {
		if end.Sub(book.UpdatedAt) < 0 {
			end = book.UpdatedAt
		}
		stat.MemoCount += book.MemoCount
	}

	if end.Sub(start) < 0 {
		stat.Duration = int64(0)
		return
	}

	stat.Duration = int64(end.Sub(start).Hours() / 24)

	return
}
