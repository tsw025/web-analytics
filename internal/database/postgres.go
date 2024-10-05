package database

import (
	"github.com/tsw025/web_analytics/internal/config"
	echologrus "github.com/tsw025/web_analytics/internal/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// we are using GORM as the ORM for the Postgres database
// This will manage the connection pool and the lifecycle of the connection
func ConnectToPostgres(cfg *config.Config) (*gorm.DB, error) {
	dsn := cfg.DatabaseURL
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		echologrus.Logger.Error("Postgres DB Connect issue")
		return nil, err
	}
	echologrus.Logger.Info("Successfully connected to Postgres DB")
	return db, nil
}
