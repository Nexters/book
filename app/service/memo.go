package service

import (
	"errors"

	"github.com/nexters/book/app/entity"
	"github.com/nexters/book/app/repository"
)

// CreateMemoParam 메모 생성 parameters
type CreateMemoParam struct {
	BookID   uint   `json:"bookId" validate:"required"`
	Text     string `json:"text" validate:"required"`
	Category string `json:"category" validate:"required"`
}

var MaxLenError = errors.New("Max character count is 150; length exceeded")

type (
	// MemoService MemoService Interface
	MemoService interface {
		FindAllMemoByUserAndBookID(userID string, bookID uint) ([]entity.Memo, error)
		CreateMemo(param CreateMemoParam, uid string) (entity.Memo, error)
	}

	// memoService struct memoService Struct
	memoService struct {
		memoRepository repository.MemoRepository
		userRepository repository.UserRepository
	}
)

// NewMemoService 생성자
func NewMemoService(mr repository.MemoRepository, ur repository.UserRepository) MemoService {
	return memoService{mr, ur}
}

// FindAllMemoByUserAndBookID userID, bookID로 해당 유저의 해당 책에 대한 모든 메모 조회
func (m memoService) FindAllMemoByUserAndBookID(userID string, bookID uint) (memos []entity.Memo, err error) {
	user, err := m.userRepository.FindUserByUID(userID)
	if err != nil {
		return
	}

	memos, err = m.memoRepository.FindAllMemoByUserAndBookID(user.ID, bookID)
	return
}

// CreateMemo 메모 생성
func (m memoService) CreateMemo(param CreateMemoParam, uid string) (memo entity.Memo, err error) {
	user, err := m.userRepository.FindUserByUID(uid)
	if err != nil {
		return
	}

	if len(memo.Text) > 150 {
		err = MaxLenError
		return
	}

	memo, err = m.memoRepository.CreateMemo(user.ID, param.BookID, param.Text, param.Category)
	return
}
