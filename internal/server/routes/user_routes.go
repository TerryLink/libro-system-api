package routes

import (
	"net/http"

	"go-api-server/internal/modules/users"

	"github.com/gin-gonic/gin"
)

// RegisterUsersRoutes 註冊書籍路由
func RegisterUsersRoutes(router *gin.Engine) {
	// 定義路由
	router.GET("/users", getAllusers)
	router.GET("/users/:id", getUserByID)
	router.POST("/users", AddNewUser)
	router.DELETE("/users/:id", deleteUser)
}

// 處理 GET /books 請求
func getAllusers(c *gin.Context) {
	allUsers := users.GetAllUsers()
	c.IndentedJSON(http.StatusOK, allUsers)
}

// 處理 GET /books/:id 請求
func getUserByID(c *gin.Context) {
	id := c.Param("id")
	user, err := users.GetUserByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

// 處理 POST /books 請求
func AddNewUser(c *gin.Context) {
	var newUser users.User
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}
	users.AddUser(newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

// 處理 PATCH /return 請求
func deleteUser(c *gin.Context) {
	id, idOk := c.GetQuery("id")
	if !idOk {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id are required"})
		return
	}
	err := users.DeleteUser(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "user deleted"})
}
