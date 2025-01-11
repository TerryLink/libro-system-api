package routes

import (
	"net/http"

	"libro-system-api/internal/database"
	"libro-system-api/internal/modules/users"

	"github.com/gin-gonic/gin"
)

// RegisterUsersRoutes 註冊書籍路由
func RegisterUsersRoutes(router *gin.Engine, db database.Service) {

	repo := users.NewUserRepository(db)
	service := users.NewUserService(repo)
	// 定義路由
	router.GET("/users", func(c *gin.Context) {
		getAllusers(c, service)
	})
	router.GET("/users/:id", func(c *gin.Context) {
		getUserByID(c, service)
	})
	router.POST("/users", func(c *gin.Context) {
		AddNewUser(c, service)
	})
	router.DELETE("/users/:id", func(c *gin.Context) {
		deleteUser(c, service)
	})
	router.PATCH("/users", func(c *gin.Context) {
		updateUser(c, service)
	})
	router.POST("/searchuser", func(c *gin.Context) {
		searchUserByEmailOrAccountName(c, service)
	})
}

// 處理 GET /books 請求
func getAllusers(c *gin.Context, service *users.UserService) {
	allUsers, err := service.FetchAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"status": "ok", "users": allUsers})
}

// 處理 GET /books/:id 請求
func getUserByID(c *gin.Context, service *users.UserService) {
	id := c.Param("id")
	user, err := service.FetchUserByID(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

// 處理 POST /books 請求
func AddNewUser(c *gin.Context, service *users.UserService) {
	var newUser users.User
	if err := c.BindJSON(&newUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}
	user, err := service.RegisterNewUser(newUser)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusCreated, user)
}

// 處理 PATCH /return 請求
func deleteUser(c *gin.Context, service *users.UserService) {
	id, idOk := c.GetQuery("id")
	if !idOk {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "id are required"})
		return

	}
	success, err := service.DeleteUser(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return

	}
	c.IndentedJSON(http.StatusOK, gin.H{"success": success})
}

func updateUser(c *gin.Context, service *users.UserService) {
	var toUpdateUser users.User
	if err := c.BindJSON(&toUpdateUser); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid request body"})
		return
	}
	user, err := service.UpdateUser(toUpdateUser)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, user)
}

func searchUserByEmailOrAccountName(c *gin.Context, service *users.UserService) {
	var req users.SearchUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	if req.AccountName == "" && req.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}
	user, err := service.SearchByEmailOrAccountName(req.AccountName, req.Email)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err.Error()})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"status": "ok", "user": user})
}
