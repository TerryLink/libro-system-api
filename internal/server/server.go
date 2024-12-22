package server

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	"go-api-server/internal/database"
	"go-api-server/internal/server/routes"
)

type Server struct {
	port int

	db     database.Service
	engine *gin.Engine
}

func NewServer(port int) *Server {
	engine := gin.Default()
	return &Server{
		port:   port,
		db:     database.New(),
		engine: engine,
	}
}

func (s *Server) Run() {
	// 註冊路由
	s.RegisterRoutes()

	// 啟動伺服器
	addr := fmt.Sprintf(":%d", s.port)
	server := &http.Server{
		Addr:         addr,
		Handler:      s.engine, // Gin Engine 作為 Handler
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	log.Printf("Server running at %s", addr)
	// 啟動伺服器，並監聽錯誤
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func (s *Server) RegisterRoutes() {
	// 假設你有多個路由需要註冊
	// 可以調用各個模組的路由註冊方法
	routes.RegisterBookRoutes(s.engine)
	routes.RegisterUsersRoutes(s.engine)

}
