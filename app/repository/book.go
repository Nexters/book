package repository

// TODO: implement
type (
	BookRepository interface{}
	bookRepository struct{}
)

func NewBookRepository() BookRepository {
	return bookRepository{}
}

func (b bookRepository) CreateBook() {}

func (b bookRepository) SearchBook() {}

func (b bookRepository) FindAllBooks() {}

func (b bookRepository) FindBookByID() {}
