package server

import (
	"context"
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
	server *http.Server
}

func NewServer(port int) *Server {
	engine := gin.Default()

	// 初始化標準 HTTP 伺服器
	addr := fmt.Sprintf(":%d", port)
	server := &http.Server{
		Addr:         addr,
		Handler:      engine, // Gin Engine 作為 Handler
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return &Server{
		port:   port,
		db:     database.New(),
		engine: engine,
		server: server,
	}
}

func (s *Server) RegisterRoutes() {
	// 註冊模組路由
	routes.RegisterBookRoutes(s.engine)
	routes.RegisterUsersRoutes(s.engine)
}

func (s *Server) Run() error {
	// 註冊路由
	s.RegisterRoutes()

	// 啟動伺服器
	log.Printf("Server running at %s", s.server.Addr)
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("Shutting down server...")
	return s.server.Shutdown(ctx)
}
