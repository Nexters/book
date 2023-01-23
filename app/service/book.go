package service

type (
	BookService interface{}
	bookService struct{}
)

func NewBookService() BookService {
	return bookService{}
}
