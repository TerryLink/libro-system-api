package books

import (
	"errors"
	"libro-system-api/internal/database"

	"gorm.io/gorm"
)

type bookRepository interface {
	GetBooks() ([]Book, error)
	GetBookByID(id string) (*Book, error)
	CheckoutBook(id string, quantity int) (*Book, error)
	ReturnBook(id string, quantity int) (*Book, error)
	AddBook(newBook Book) error
}

// BookRepository
type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db database.Service) *BookRepository {
	return &BookRepository{db: db.GetDB()}
}

// GetBooks get all books
func (r *BookRepository) GetBooks() ([]Book, error) {
	var books []Book
	if err := r.db.Find(&books).Error; err != nil {
		return nil, err
	}
	return books, nil
}

// GetBookByID
func (r *BookRepository) GetBookByID(id string) (*Book, error) {

	var book Book
	if err := r.db.First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("book not found")
		}
		return nil, err
	}
	return &book, nil
}

// Search book by title or author
func (r *BookRepository) SearchBooks(keyword string) ([]Book, error) {
	var books []Book
	query := "%" + keyword + "%"
	if err := r.db.Where("title Like ? or author Like ?", query, query).Find(&books).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("book not found")
		}
		return nil, err
	}
	return books, nil
}

// rent book
func (r *BookRepository) CheckoutBook(id string, quantity int) (*Book, error) {
	var book Book
	if err := r.db.First(&book, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("book not found")
		}
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
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("book not found")
		}
		return nil, err
	}
	book.Quantity += quantity
	if err := r.db.Save(&book).Error; err != nil {
		return nil, err
	}
	return &book, nil
}

// AddBook add new book
func (r *BookRepository) AddBook(newBook Book) error {
	// check newBook info
	if newBook.Title == "" || newBook.Author == "" {
		return errors.New("title or author are required")
	}
	if err := r.db.Create(&newBook).Error; err != nil {
		return err
	}
	return nil
}
