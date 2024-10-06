package analyze

import (
	"github.com/labstack/echo/v4"
)

func (h *AnalyseHandler) RegisterRoutes(e *echo.Group, mw echo.MiddlewareFunc) {
	analyseGroup := e.Group("/analyse", mw)
	analyseGroup.POST("", h.Analyse)
}
