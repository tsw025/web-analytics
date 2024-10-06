package analyze

import (
	"github.com/labstack/echo/v4"
	"github.com/tsw025/web_analytics/internal/models"
	"github.com/tsw025/web_analytics/internal/repositories"
	"github.com/tsw025/web_analytics/internal/schemas"
	"github.com/tsw025/web_analytics/internal/services"
)

type AnalyseHandler struct {
	analyzeService services.AnalyseService
	userRpo        repositories.UserRepository
}

func NewAnalyseHandler(
	analyzeService services.AnalyseService,
	userRepo repositories.UserRepository,
) *AnalyseHandler {
	return &AnalyseHandler{
		analyzeService: analyzeService,
		userRpo:        userRepo,
	}
}

func (h *AnalyseHandler) Analyse(c echo.Context) error {
	req := new(schemas.AnalyserRequest)
	if err := c.Bind(req); err != nil {
		return err
	}
	dbUser, _ := c.Get("user").(*models.User)
	if err := c.Validate(req); err != nil {
		return err
	}

	dbAnalyse, err := h.analyzeService.Analyse(req.URL, dbUser)
	if err != nil {
		return err
	}

	return c.JSON(200, dbAnalyse)
}
