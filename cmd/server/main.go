package main

import (
	"github.com/go-playground/validator/v10"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/tsw025/web_analytics/internal/config"
	middleware2 "github.com/tsw025/web_analytics/internal/custom_middleware"
	"github.com/tsw025/web_analytics/internal/database"
	"github.com/tsw025/web_analytics/internal/echologrus"
	"github.com/tsw025/web_analytics/internal/handlers"
	"github.com/tsw025/web_analytics/internal/handlers/analyze"
	"github.com/tsw025/web_analytics/internal/handlers/auth"
	"github.com/tsw025/web_analytics/internal/logger"
	"github.com/tsw025/web_analytics/internal/repositories"
	"github.com/tsw025/web_analytics/internal/schemas"
	"github.com/tsw025/web_analytics/internal/services"
	"time"

	"github.com/labstack/echo/v4"
)

func main() {
	// Echo instance
	e := echo.New()

	// Config
	cfg, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}
	//Logger
	logger.InitLogger(cfg, e)

	// We are setting the timezone to UTC, so that all the time values are stored in UTC
	time.Local = time.UTC

	// Initialize the database
	db, err := database.ConnectToPostgres(cfg)
	if err != nil {
		panic(err)
	}

	// JwtMiddleware
	jwtMiddleware := echojwt.WithConfig(echojwt.Config{
		SigningKey:     []byte(cfg.JWTSecret),
		SuccessHandler: middleware2.JWTSuccessHandler(repositories.NewUserRepository(db)),
	})
	// Middleware
	e.Use(middleware.Recover()) // Recover from panics anywhere in the middleware chain
	e.HTTPErrorHandler = handlers.CustomHTTPErrorHandler

	// Validator
	e.Validator = &schemas.BaseValidator{Validator: validator.New()}

	// Service Initialization
	authService := services.NewPasswordAuthService(repositories.NewUserRepository(db))
	tokenService := services.NewAuthTokenService(cfg)
	analyseService := services.NewAnalyseService(repositories.NewWebsiteRepository(db))

	// Handler Initialization
	mainGroup := e.Group("/api")
	authHandler := auth.NewAuthHandler(authService, tokenService)
	authHandler.RegisterRoutes(mainGroup)

	analyzeHandler := analyze.NewAnalyseHandler(analyseService, repositories.NewUserRepository(db))
	analyzeHandler.RegisterRoutes(mainGroup, jwtMiddleware)

	echologrus.Logger.Debugf("Starting server on port %s", cfg.ServerPort)
	e.Logger.Fatal(e.Start(":" + cfg.ServerPort))
}
