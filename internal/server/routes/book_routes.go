package routes

import (
	"net/http"
	"strconv"

	"go-api-server/internal/modules/books"

	"github.com/gin-gonic/gin"
)

// RegisterBookRoutes 註冊書籍路由
func RegisterBookRoutes(router *gin.Engine) {
	// 定義路由
	router.GET("/books", getBooks)
	router.GET("/books/:id", getBookByID)
	router.POST("/books", createBook)
	router.PATCH("/checkout", checkoutBook)
	router.PATCH("/return", returnBook)
}

// 處理 GET /books 請求
func getBooks(c *gin.Context) {
	allBooks := books.FetchAllBooks()
	c.JSON(http.StatusOK, gin.H{"status": "ok", "books": allBooks})
}

// 處理 GET /books/:id 請求
func getBookByID(c *gin.Context) {
	id := c.Param("id")
	book, err := books.FetchBookByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

// 處理 POST /books 請求
func createBook(c *gin.Context) {
	var newBook books.Book
	if err := c.BindJSON(&newBook); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}
	books.CreateNewBook(newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

// 處理 PATCH /checkout 請求
func checkoutBook(c *gin.Context) {
	id, idOk := c.GetQuery("id")
	qun, qunOk := c.GetQuery("quantity")
	if !idOk || !qunOk {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id and quantity are required"})
		return
	}
	quantity, err := strconv.Atoi(qun)
	if err != nil || quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "quantity must be a valid positive integer"})
		return
	}
	book, err := books.CheckoutBook(id, quantity)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}

// 處理 PATCH /return 請求
func returnBook(c *gin.Context) {
	id, idOk := c.GetQuery("id")
	qun, qunOk := c.GetQuery("quantity")
	if !idOk || !qunOk {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id and quantity are required"})
		return
	}
	quantity, err := strconv.Atoi(qun)
	if err != nil || quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "quantity must be a valid positive integer"})
		return
	}
	book, err := books.ReturnBook(id, quantity)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, book)
}
