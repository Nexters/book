package repository

type (
	MemoRepository interface{}
	memoRepository struct{}
)

func NewMemoRepository() MemoRepository {
	return memoRepository{}
}

func (m memoRepository) CreateMemo() {}

func (m memoRepository) FindAllMemoByBookID() {}
