package repository

import (
	"errors"

	"github.com/nexters/book/app/config"
	"github.com/nexters/book/app/entity"
)

type (
	// MemoRepository MemoRepository Interface
	MemoRepository interface {
		FindAllMemoByUserAndBookID(userID uint, bookID uint) ([]entity.Memo, error)
		CreateMemo(userID uint, bookID uint, text string, category string) (entity.Memo, error)
	}

	// memoRepository memoRepository Struct
	memoRepository struct {
		db config.Database
	}
)

// NewMemoRepository 생성자
func NewMemoRepository(db config.Database) MemoRepository {
	return memoRepository{db}
}

// CreateMemo 메모 생성
func (m memoRepository) CreateMemo(
	userID uint,
	bookID uint,
	text string,
	category string,
) (memo entity.Memo, err error) {
	memo = entity.Memo{
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

// FindAllMemoByUserAndBookID 메모 조회
func (m memoRepository) FindAllMemoByUserAndBookID(userID uint, bookID uint) (memos []entity.Memo, err error) {
	memos = []entity.Memo{}

	if res := m.db.Where("book_id = ? AND user_id = ?", bookID, userID).Find(&memos); res.Error != nil {
		err = res.Error
		return
	}

	return
}
