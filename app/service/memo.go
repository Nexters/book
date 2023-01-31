package service

import (
	"github.com/nexters/book/app/entity"
	"github.com/nexters/book/app/repository"
)

type CreateMemoParam struct {
	BookID   uint   `json:"bookId" validate:"required"`
	Text     string `json:"text" validate:"required"`
	Category string `json:"category" validate:"required"`
}

type (
	MemoService interface {
		FindAllMemoByUserAndBookID(userID string, bookID uint) ([]entity.Memo, error)
		CreateMemo(param CreateMemoParam, uid string) (entity.Memo, error)
	}
	memoService struct {
		memoRepository repository.MemoRepository
		userRepository repository.UserRepository
	}
)

func NewMemoService(mr repository.MemoRepository, ur repository.UserRepository) MemoService {
	return memoService{mr, ur}
}

func (m memoService) FindAllMemoByUserAndBookID(userID string, bookID uint) (memos []entity.Memo, err error) {
	user, err := m.userRepository.FindUserByUID(userID)
	if err != nil {
		return
	}

	memos, err = m.memoRepository.FindAllMemoByUserAndBookID(user.ID, bookID)
	return
}

func (m memoService) CreateMemo(param CreateMemoParam, uid string) (memo entity.Memo, err error) {
	user, err := m.userRepository.FindUserByUID(uid)
	if err != nil {
		return
	}

	memo, err = m.memoRepository.CreateMemo(user.ID, param.BookID, param.Text, param.Category)
	return
}
