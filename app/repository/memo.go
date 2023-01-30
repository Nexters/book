package repository

import (
	"errors"

	"github.com/nexters/book/app/config"
	"github.com/nexters/book/app/entity"
)

type (
	MemoRepository interface {
		FindAllMemoByUserAndBookID(userID uint, bookID uint) ([]entity.Memo, error)
		CreateMemo(userID uint, bookID uint, text string, category string) (entity.Memo, error)
	}
	memoRepository struct {
		db config.Database
	}
)

func NewMemoRepository(db config.Database) MemoRepository {
	return memoRepository{db}
}

func (m memoRepository) CreateMemo(
	userID uint,
	bookID uint,
	text string,
	category string,
) (memo entity.Memo, err error) {
	memo = entity.Memo{
		UserID:   userID,
		BookID:   bookID,
		Text:     text,
		Category: category,
	}

	res := m.db.Create(&memo)
	if res.Error != nil {
		err = res.Error
		return
	}

	if res.RowsAffected == 0 {
		err = errors.New("Memo creation failed")
	}
	return
}

func (m memoRepository) FindAllMemoByUserAndBookID(userID uint, bookID uint) (memos []entity.Memo, err error) {
	memos = []entity.Memo{}

	if res := m.db.Where("book_id = ? AND user_id = ?", bookID, userID).Find(&memos); res.Error != nil {
		err = res.Error
		return
	}

	return
}
