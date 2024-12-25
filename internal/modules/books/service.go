package books

type BookService struct {
	repo *BookRepository
}

func NewBookService(repo *BookRepository) *BookService {
	return &BookService{repo: repo}
}

// FetchAllBooks 獲取所有書籍
func (s *BookService) FetchAllBooks() ([]Book, error) {
	return s.repo.GetBooks()
	// return GetBooks()
}

// FetchBookByID 根據 ID 獲取書籍
func (s *BookService) FetchBookByID(id string) (*Book, error) {
	return s.repo.GetBookByID(id)
	// return GetBookByID(id)
}

// CheckoutBook 借閱書籍
func (s *BookService) CheckoutBook(id string, quantity int) (*Book, error) {
	return s.repo.CheckoutBook(id, quantity)
	// book, err := GetBookByID(id)
	// if err != nil {
	// 	return nil, err
	// }
	// if book.Quantity < quantity {
	// 	return nil, errors.New("insufficient stock")
	// }
	// book.Quantity -= quantity
	// return book, nil
}

// ReturnBook 歸還書籍
func (s *BookService) ReturnBook(id string, quantity int) (*Book, error) {
	return s.repo.ReturnBook(id, quantity)
}

// CreateNewBook 新增書籍
func (s *BookService) CreateNewBook(newBook Book) {
	// AddBook(newBook)
	s.repo.AddBook(newBook)
}
