package memo

import (
	"errors"

	"github.com/nexters/book/app/entity"
	"github.com/nexters/book/app/user"
)

// CreateMemoParam 메모 생성 parameters
type CreateMemoParam struct {
	BookID   uint   `json:"bookId" validate:"required"`
	Text     string `json:"text" validate:"required"`
	Category string `json:"category" validate:"required"`
}

// UpdateMemoParam 메모 생성 parameters
type UpdateMemoParam struct {
	MemoID   uint   `json:"bookId" validate:"required"`
	Text     string `json:"text"`
	Category string `json:"category"`
}

var MaxLenError = errors.New("Max character count is 150; length exceeded")

type (
	// MemoService MemoService Interface
	MemoService interface {
		FindAllMemoByUserAndBookID(userID string, bookID uint) ([]entity.Memo, error)
		CreateMemo(param CreateMemoParam, uid string) (entity.Memo, error)
		UpdateMemo(param UpdateMemoParam) (entity.Memo, error)
		DeleteMemo(memoID uint) (entity.Memo, error)
	}

	// memoService struct memoService Struct
	memoService struct {
		memoRepository MemoRepository
		userRepository user.UserRepository
	}
)

// NewMemoService 생성자
func NewMemoService(mr MemoRepository, ur user.UserRepository) MemoService {
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
	if err != nil {
		return
	}

	if len(memo.Text) > 150 {
		err = MaxLenError
		return
	}

	memo, err = m.memoRepository.CreateMemo(param.BookID, param.Text, param.Category)
	return
}

// UpdateMemo 메모 업데이트
func (m memoService) UpdateMemo(param UpdateMemoParam) (memo entity.Memo, err error) {
	return m.memoRepository.UpdateMemo(param.MemoID, param.Text, param.Category)
}

// DeleteMemo 메모 삭제
func (m memoService) DeleteMemo(memoID uint) (memo entity.Memo, err error) {
	return m.memoRepository.DeleteMemo(memoID)
}
