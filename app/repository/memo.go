package repository

import (
	"errors"
	"time"

	"github.com/nexters/book/app/entity"
	"github.com/nexters/book/config"
)

type (
	// MemoRepository MemoRepository Interface
	MemoRepository interface {
		FindAllMemoByUserAndBookID(userID uint, bookID uint) ([]entity.Memo, error)
		CreateMemo(bookID uint, text string, category string) (entity.Memo, error)
		UpdateMemo(memoID uint, text string, category string) (entity.Memo, error)
		DeleteMemo(memoID uint) (entity.Memo, error)
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
	bookID uint,
	text string,
	category string,
) (memo entity.Memo, err error) {
	memo = entity.Memo{
		Text:     text,
		Category: category,
		BookID:   bookID,
	}

	// transaction 시작
	tx := m.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 메모 추가
	if err = tx.Create(&memo).Error; err != nil {
		tx.Rollback()
		return
	}

	// book의 updated_at 갱신
	now := time.Now()
	if err = tx.Model(&entity.Book{}).Where("books.id", bookID).Update("updated_at", now).Error; err != nil {
		tx.Rollback()
		return
	}

	// transaction 종료
	err = tx.Commit().Error

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

// UpdateMemo
func (m memoRepository) UpdateMemo(memoID uint, text string, category string) (memo entity.Memo, err error) {
	memo.ID = memoID
	if len(text) > 0 {
		memo.Text = text
	}
	if len(category) > 0 {
		memo.Category = category
	}

	tx := m.db.Model(&memo).Updates(memo)

	if err = tx.Error; err != nil {
		return
	}

	if tx.RowsAffected == 0 {
		err = errors.New("Memo update failed")
	}

	return

}

// DeleteMemo
func (m memoRepository) DeleteMemo(memoID uint) (memo entity.Memo, err error) {
	memo.ID = memoID
	tx := m.db.Delete(&memo)

	if err = tx.Error; err != nil {
		return
	}

	if tx.RowsAffected == 0 {
		err = errors.New("Memo delete failed")
		return
	}

	return
}
