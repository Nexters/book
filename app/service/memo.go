package service

import (
	"github.com/nexters/book/app/entity"
	"github.com/nexters/book/app/repository"
)

type CreateMemoParam struct {
	UserID   string `json:"userId"`
	BookID   uint   `json:"bookId"`
	Text     string `json:"text"`
	Category string `json:"category"`
}

type (
	MemoService interface {
		FindAllMemoByUserAndBookID(userID string, bookID uint) ([]entity.Memo, error)
		CreateMemo(param CreateMemoParam) (entity.Memo, error)
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
	// TODO: user repository -> find user by uid
	user, err := m.userRepository.FindUserByUID(userID)
	if err != nil {
		return
	}
	memos, err = m.memoRepository.FindAllMemoByUserAndBookID(user.ID, bookID)
	return
}

func (m memoService) CreateMemo(param CreateMemoParam) (memo entity.Memo, err error) {
	user, err := m.userRepository.FindUserByUID(param.UserID)
	if err != nil {
		return
	}

	memo, err = m.memoRepository.CreateMemo(user.ID, param.BookID, param.Text, param.Category)
	return
}
