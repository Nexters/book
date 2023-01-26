package repository

import (
	"github.com/nexters/book/app/config"
	"github.com/nexters/book/app/entity"
)

type (
	MemoRepository interface {
		FindAllMemoByUserAndBookID(userID uint, bookID uint) ([]entity.Memo, error)
	}
	memoRepository struct {
		db config.Database
	}
)

func NewMemoRepository(db config.Database) MemoRepository {
	return memoRepository{db}
}

func (m memoRepository) CreateMemo() {}

func (m memoRepository) FindAllMemoByUserAndBookID(userID uint, bookID uint) (memos []entity.Memo, err error) {
	memos = []entity.Memo{}

	if res := m.db.Where("book_id ? AND user_id = ?", bookID, userID); res.Error != nil {
		err = res.Error
		return
	}

	return
}
