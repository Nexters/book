package service

import "github.com/nexters/book/app/entity"

type (
	MemoService interface {
		FindAllMemoByUserAndBookID(userID string, bookID uint64) ([]entity.Memo, error)
		CreateMemo(param interface{}) (entity.Memo, error)
	}
	memoService struct{}
)

func NewMemoService() MemoService {
	return memoService{}
}

func (m memoService) FindAllMemoByUserAndBookID(userID string, bookID uint64) ([]entity.Memo, error) {
	memos := []entity.Memo{}

	return memos, nil
}

func (m memoService) CreateMemo(param interface{}) (entity.Memo, error) {
	memo := entity.Memo{}

	return memo, nil
}
