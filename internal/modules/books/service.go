package books

type BookService struct {
	repo *BookRepository
}

func NewBookService(repo *BookRepository) *BookService {
	return &BookService{repo: repo}
}

// FetchAllBooks
func (s *BookService) FetchAllBooks() ([]Book, error) {
	return s.repo.GetBooks()
	// return GetBooks()
}

// FetchBookByID
func (s *BookService) FetchBookByID(id string) (*Book, error) {
	return s.repo.GetBookByID(id)
	// return GetBookByID(id)
}

// CheckoutBook
func (s *BookService) CheckoutBook(id string, quantity int) (*Book, error) {
	return s.repo.CheckoutBook(id, quantity)
}

func (s *BookService) SearchBooks(keyword string) ([]Book, error) {
	return s.repo.SearchBooks(keyword)
}

// ReturnBook
func (s *BookService) ReturnBook(id string, quantity int) (*Book, error) {
	return s.repo.ReturnBook(id, quantity)
}

// CreateNewBook
func (s *BookService) CreateNewBook(newBook Book) {
	// AddBook(newBook)
	s.repo.AddBook(newBook)
}
