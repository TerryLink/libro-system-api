package routes

import (
	"net/http"
	"strconv"

	"libro-system-api/internal/modules/books"

	"libro-system-api/internal/database"

	"github.com/gin-gonic/gin"
)

// RegisterBookRoutes
func RegisterBookRoutes(router *gin.Engine, db database.Service) {
	// defined routes

	repo := books.NewBookRepository(db)
	service := books.NewBookService(repo)

	router.GET("/books", func(c *gin.Context) {
		book, err := service.FetchAllBooks()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok", "books": book})
	})
	router.GET("/books/:id", func(c *gin.Context) {
		id := c.Param("id")
		book, err := service.FetchBookByID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"message": "book not found"})
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok", "book": book})
	})
	router.GET("/books", func(c *gin.Context) {
		keyword, keywordOK := c.GetQuery("keyword")
		if !keywordOK {
			c.JSON(http.StatusBadRequest, gin.H{"message": "keyword is required"})
		}
		books, err := service.SearchBooks(keyword)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok", "books": books})
	})
	router.POST("/books", func(c *gin.Context) {
		var newBook books.Book
		if err := c.BindJSON(&newBook); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		}
		service.CreateNewBook(newBook)
		c.JSON(http.StatusCreated, gin.H{"status": "ok", "book": newBook})
	})
	router.PATCH("/checkout", func(c *gin.Context) {
		id, idOK := c.GetQuery("id")
		qun, qunOK := c.GetQuery("quantity")
		if !idOK || !qunOK {
			c.JSON(http.StatusBadRequest, gin.H{"message": "id and quantity are required"})
		}
		quantity, err := strconv.Atoi(qun)
		if err != nil || quantity < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "quantity must be a valid positive integer"})
		}
		book, err := service.CheckoutBook(id, quantity)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok", "book": book})
	})
	router.PATCH("/return", func(c *gin.Context) {
		id, idOK := c.GetQuery("id")
		qun, qunOK := c.GetQuery("quantity")
		if !idOK || !qunOK {
			c.JSON(http.StatusBadRequest, gin.H{"message": "id and quantity are required"})
		}
		quantity, err := strconv.Atoi(qun)
		if err != nil || quantity < 0 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "quantity must be a valid positive integer"})
		}
		book, err := service.ReturnBook(id, quantity)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		}
		c.JSON(http.StatusOK, gin.H{"status": "ok", "book": book})
	})
}
