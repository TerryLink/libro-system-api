package books

import (
	"errors"
	"go-api-server/internal/database"

	"gorm.io/gorm"
)

// BookRepository 書籍資料庫操作接口
type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db database.Service) *BookRepository {
	return &BookRepository{db: db.GetDB()}
}

// 模擬數據庫
// var books = []Book{
// 	{ID: "1", Title: "Golang", Author: "Google", Quantity: 10},
// 	{ID: "2", Title: "Java", Author: "Oracle", Quantity: 20},
// 	{ID: "3", Title: "Python", Author: "Python Software Foundation", Quantity: 30},
// }

// GetBooks 獲取所有書籍
func (r *BookRepository) GetBooks() ([]Book, error) {
	var books []Book
	if err := r.db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

// GetBookByID 根據 ID 獲取書籍
func (r *BookRepository) GetBookByID(id string) (*Book, error) {

	var book Book
	if err := r.db.First(&book, id).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

func (r *BookRepository) CheckoutBook(id string, quantity int) (*Book, error) {
	var book Book
	if err := r.db.First(&book, id).Error; err != nil {
		return nil, err
	}
	if book.Quantity < quantity {
		return nil, errors.New("insufficient stock")
	}
	book.Quantity -= quantity
	if err := r.db.Save(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil

}

func (r *BookRepository) ReturnBook(id string, quantity int) (*Book, error) {
	var book Book
	if err := r.db.First(&book, id).Error; err != nil {
		return nil, err
	}
	book.Quantity += quantity
	if err := r.db.Save(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

// AddBook 新增書籍
func (r *BookRepository) AddBook(newBook Book) error {
	if err := r.db.Create(&newBook).Error; err != nil {
		return err
	}
	return nil
}
