package websites

import (
	"github.com/labstack/echo/v4"
	"github.com/tsw025/web_analytics/internal/models"
	"github.com/tsw025/web_analytics/internal/repositories"
	"github.com/tsw025/web_analytics/internal/services"
	"strconv"
)

type WebsiteHandler struct {
	websiteService services.WebsiteService
	userRepo       repositories.UserRepository
	websiteRepo    repositories.WebsiteRepository
}

func NewWebsiteHandler(
	websiteService services.WebsiteService,
	userRepo repositories.UserRepository,
	websiteRepo repositories.WebsiteRepository,
) *WebsiteHandler {
	return &WebsiteHandler{
		websiteService: websiteService,
		userRepo:       userRepo,
		websiteRepo:    websiteRepo,
	}
}

func (h *WebsiteHandler) GetWebsites(c echo.Context) error {
	dbUser, _ := c.Get("user").(*models.User)

	websites, err := h.websiteService.GetWebsites(dbUser)
	if err != nil {
		return err
	}

	return c.JSON(200, websites)
}

func (h *WebsiteHandler) GetWebsite(c echo.Context) error {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)

	if err != nil {
		return err
	}

	website, err := h.websiteService.GetWebsite(uint(intId))
	if err != nil {
		return err
	}

	return c.JSON(200, website)
}

func (h *WebsiteHandler) UpdateWebsite(c echo.Context) error {
	id := c.Param("id")
	intId, err := strconv.Atoi(id)

	if err != nil {
		return err
	}

	req := new(models.Website)
	if err := c.Bind(req); err != nil {
		return err
	}

	website, err := h.websiteService.UpdateWebsite(uint(intId), req)
	if err != nil {
		return err
	}

	return c.JSON(200, website)
}
