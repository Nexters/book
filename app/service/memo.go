package service

import (
	"errors"

	"github.com/nexters/book/app/config"
	"github.com/nexters/book/app/entity"
)

type CreateMemoParam struct {
	UserID   string `json:"userId"`
	BookID   uint64 `json:"bookId"`
	Text     string `json:"text"`
	Category string `json:"category"`
}

type (
	MemoService interface {
		FindAllMemoByUserAndBookID(userID string, bookID uint64) ([]entity.Memo, error)
		CreateMemo(param CreateMemoParam) (entity.Memo, error)
	}
	memoService struct {
		db config.Database
	}
)

func NewMemoService(db config.Database) MemoService {
	return memoService{db}
}

func (m memoService) FindAllMemoByUserAndBookID(userID string, bookID uint64) ([]entity.Memo, error) {
	memos := []entity.Memo{}

	user := entity.User{}
	res := m.db.Where("uid = ?", userID).First(&user)

	if res.Error != nil || res.RowsAffected == 0 {
		return memos, errors.New("User not found")
	}

	res = m.db.Where("book_id = ? AND user_id = ?", bookID, user.ID).Find(&memos)
	if res.Error != nil || res.RowsAffected == 0 {
		return memos, errors.New("Memo not found")
	}

	return memos, nil
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
