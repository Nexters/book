package service

import (
	"errors"

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

func (m memoService) CreateMemo(param CreateMemoParam) (entity.Memo, error) {
	user := entity.User{}
	memo := entity.Memo{
		BookID:   param.BookID,
		Text:     param.Text,
		Category: param.Category,
	}
	res := m.db.Where("uid = ?", param.UserID).First(&user)
	if res.Error != nil {
		return memo, res.Error
	}

	if res.RowsAffected == 0 {
		return memo, errors.New("User not found")
	}

	memo.UserID = uint64(user.ID)

	res = m.db.Create(&memo)
	if res.Error != nil || res.RowsAffected == 0 {
		return memo, errors.New("Memo creation failed")
	}

	return memo, nil
}
