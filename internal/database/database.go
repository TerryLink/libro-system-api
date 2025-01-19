package database

import (
	"context"
	"fmt"
	"libro-system-api/internal/config"
	"log"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Service interface {
	Health() map[string]string
	Close() error
	GetDB() *gorm.DB
}

// GetDB returns the underlying GORM DB instance.
func (s *service) GetDB() *gorm.DB {
	return s.db
}

type service struct {
	db *gorm.DB
}

var (
	dbInstance *service
)

// New initializes a new GORM service instance.
func New() Service {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}
	cfg, err := config.LoadConfig()
	if err != nil {
		if err != nil {
			log.Fatalf("Failed to load configuration: %v", err)
		}
	}
	// GORM Logger (optional for debugging)
	newLogger := logger.Default.LogMode(logger.Info)

	// Create DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", cfg.Database.User, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)

	// Open database using GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Get the underlying sql.DB connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get sql.DB from GORM: %v", err)
	}

	// Set connection pool parameters
	sqlDB.SetConnMaxLifetime(0)
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(50)

	dbInstance = &service{
		db: db,
	}
	return dbInstance
}

// Health checks the health of the database connection by pinging the database.
func (s *service) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Get the underlying sql.DB from GORM
	sqlDB, err := s.db.DB()
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("failed to get sql.DB: %v", err)
		log.Fatalf("failed to get sql.DB: %v", err)
		return stats
	}

	// Ping the database
	err = sqlDB.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		log.Fatalf("db down: %v", err)
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Get database stats
	dbStats := sqlDB.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	return stats
}

// Close closes the database connection.
func (s *service) Close() error {
	sqlDB, err := s.db.DB()
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to get sql.DB for closing: %v", err)
	}
	log.Printf("Disconnected from database: %s", cfg.Database.Name)
	return sqlDB.Close()
}
