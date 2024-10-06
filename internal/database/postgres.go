package database

import (
	"github.com/tsw025/web_analytics/internal/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// we are using GORM as the ORM for the Postgres database
// This will manage the connection pool and the lifecycle of the connection
func connectToPostgres(cfg *config.Config) (*gorm.DB, error) {
	dsn := cfg.DatabaseURL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
