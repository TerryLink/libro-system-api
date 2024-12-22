package books

import "errors"

// FetchAllBooks 獲取所有書籍
func FetchAllBooks() []Book {
	return GetBooks()
}

// FetchBookByID 根據 ID 獲取書籍
func FetchBookByID(id string) (*Book, error) {
	return GetBookByID(id)
}

// CheckoutBook 借閱書籍
func CheckoutBook(id string, quantity int) (*Book, error) {
	book, err := GetBookByID(id)
	if err != nil {
		return nil, err
	}
	if book.Quantity < quantity {
		return nil, errors.New("insufficient stock")
	}
	book.Quantity -= quantity
	return book, nil
}

// ReturnBook 歸還書籍
func ReturnBook(id string, quantity int) (*Book, error) {
	book, err := GetBookByID(id)
	if err != nil {
		return nil, err
	}
	book.Quantity += quantity
	return book, nil
}

// CreateNewBook 新增書籍
func CreateNewBook(newBook Book) {
	AddBook(newBook)
}
