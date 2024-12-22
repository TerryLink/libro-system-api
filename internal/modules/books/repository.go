package books

import (
	"errors"
)

// 模擬數據庫
var books = []Book{
	{ID: "1", Title: "Golang", Author: "Google", Quantity: 10},
	{ID: "2", Title: "Java", Author: "Oracle", Quantity: 20},
	{ID: "3", Title: "Python", Author: "Python Software Foundation", Quantity: 30},
}

// GetBooks 獲取所有書籍
func GetBooks() []Book {
	return books
}

// GetBookByID 根據 ID 獲取書籍
func GetBookByID(id string) (*Book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}
	return nil, errors.New("book not found")
}

// AddBook 新增書籍
func AddBook(newBook Book) {
	books = append(books, newBook)
}
