package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"

	"libro-system-api/internal/database"
	"libro-system-api/internal/server/routes"
)

type Server struct {
	port int

	db     database.Service
	engine *gin.Engine
	server *http.Server
}

func NewServer(port int) *Server {
	engine := gin.Default()

	// init standard HTTP server
	addr := fmt.Sprintf(":%d", port)
	server := &http.Server{
		Addr:         addr,
		Handler:      engine, // Gin Engine as Handler
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
	// register models routes
	routes.RegisterBookRoutes(s.engine, s.db)
	routes.RegisterUsersRoutes(s.engine, s.db)
}

func (s *Server) Run() error {
	// register routes
	s.RegisterRoutes()

	// start server
	log.Printf("Server running at %s", s.server.Addr)
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Println("Shutting down server...")
	return s.server.Shutdown(ctx)
}
