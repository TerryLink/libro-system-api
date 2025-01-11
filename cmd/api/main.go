package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"libro-system-api/internal/server"
)

func gracefulShutdown(apiServer *server.Server, done chan bool) {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// Listen for the interrupt signal.
	<-ctx.Done()

	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := apiServer.Shutdown(ctx); err != nil {
		log.Printf("Server forced to shutdown with error: %v", err)
	}

	log.Println("Server exiting")

	// Notify the main goroutine that the shutdown is complete
	done <- true
}

func main() {
	// 初始化伺服器
	apiServer := server.NewServer(8080)

	// 創建一個 channel，用於通知關閉完成
	done := make(chan bool, 1)

	// 啟動優雅關閉的 goroutine
	go gracefulShutdown(apiServer, done)

	// 啟動伺服器
	if err := apiServer.Run(); err != nil && err != http.ErrServerClosed {
		panic(fmt.Sprintf("http server error: %s", err))
	}

	// 等待優雅關閉完成
	<-done
	log.Println("Graceful shutdown complete.")
}
