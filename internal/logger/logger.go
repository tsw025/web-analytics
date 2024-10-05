package logger

import (
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"github.com/tsw025/web_analytics/internal/config"
	echologrus "github.com/tsw025/web_analytics/internal/middleware"
	"time"
)

var Logger *logrus.Logger

func InitLogger(cfg *config.Config, e *echo.Echo) {
	Logger = logrus.New()
	Logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: time.RFC3339,
		FullTimestamp:   true,
	})
	Logger.SetReportCaller(true)
	Logger.SetLevel(cfg.LogLevel)
	if cfg.LogLevel == logrus.DebugLevel {
		Logger.Info("Debug mode enabled")
		e.Debug = true
	}

	echologrus.Logger = Logger
	e.Logger = echologrus.GetEchoLogger()
	e.Use(echologrus.Hook())
}
